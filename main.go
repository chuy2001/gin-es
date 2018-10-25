package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/chuy2001/gin-es/config"
	"github.com/chuy2001/gin-es/controllers"
	"github.com/chuy2001/gin-es/db"
	"github.com/chuy2001/gin-es/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/olivere/elastic.v5"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func ContextInjector(ESClient *elastic.Client, AMQPChannel string, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("ESClient", ESClient)
		c.Set("AMQPChannel", AMQPChannel)
		c.Set("Config", config)

		c.Next()
	}
}

//Init elastic
func initESClient(url string, indices []string, doSniff bool) *elastic.Client {

	log.Printf("Connecting to ES on: %v", url)
	elasticClient, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(doSniff))
	models.CheckFatalError(err)

	log.Println("Connected to ES")

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := elasticClient.Ping("http://127.0.0.1:9200").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := elasticClient.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	for _, index := range indices {

		log.Printf("Initializing Index: %s", index)

		indexExists, err := elasticClient.IndexExists(index).Do(context.Background())
		models.CheckFatalError(err)
		if !indexExists {
			resp, err := elasticClient.CreateIndex(index).Do(context.Background())
			models.CheckFatalError(err)
			if !resp.Acknowledged {
				log.Fatalf("Cannot create index: %s on ES", index)
			}
			log.Printf("Created index: %s on ES", index)

		} else {
			log.Printf("Index: %s already exists on ES", index)
		}

		_, err = elasticClient.OpenIndex(index).Do(context.Background())
		models.CheckFatalError(err)

		mapping, err := elasticClient.GetMapping().Index(index).Do(context.Background())
		if err != nil {
			log.Printf("Cannot get mapping for index: %s", index)
		}
		log.Printf("Mapping for index %s: %s", index, mapping)
	}

	return elasticClient
}
func main() {
	r := gin.Default()

	// store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	store, _ := sessions.NewRedisStore(10, "tcp", "192.168.1.123:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("gin-boilerplate-session", store))

	r.Use(CORSMiddleware())

	log.Println("Starting elasticgin") 
	//Config fetch
	config := config.GetConfFromJSONFile("config.json")
	//ES init
	esClient := initESClient(config.ElasticURL, config.Indices, config.SniffCluster)
	defer esClient.Stop()
	r.Use(ContextInjector(esClient, "null", config))

	db.Init()

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/user/signin", user.Signin)
		v1.POST("/user/signup", user.Signup)
		v1.GET("/user/signout", user.Signout)

		/*** START ElasticSearch ***/
		es := new(controllers.EsController)
		v1.POST("/_search", es.Search)
		v1.POST("/_msearch", es.MSearch)
	}

	r.LoadHTMLGlob("./public/html/*")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"ECharts入门示例 - 柱状图":      "bar",
			"ECharts入门示例 - rest-get": "rest",
			"ECharts入门示例 - es-post":  "es1",
		})
	})

	r.GET("/bar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "bar.html", gin.H{})
	})
	r.GET("/rest", func(c *gin.Context) {
		c.HTML(http.StatusOK, "rest.html", gin.H{})
	})
	r.GET("/es1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "es1.html", gin.H{})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run(":9000")
}
