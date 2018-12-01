package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chuy2001/gin-es/config"
	"github.com/chuy2001/gin-es/controllers"
	"github.com/chuy2001/gin-es/db"

	// "github.com/gin-gonic/contrib/sessions"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	// "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
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

func ContextInjector( AMQPChannel string, config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("AMQPChannel", AMQPChannel)
		c.Set("Config", config)

		c.Next()
	}
}

func main() {
	config := config.GetConfFromJSONFile("config.json")
	r := gin.Default()

	// store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	// store, _ := sessions.NewRedisStore(10, "tcp", "192.168.1.123:6379", "", []byte("secret"))
	// store, _ := redis.NewRedisStore(10, "tcp", "192.168.1.123:6379", "", []byte("secret"))
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("gin-boilerplate-session", store))
	r.Use(CORSMiddleware())

	//Config fetch
	log.Println("Starting sql & elasticsearch") 
	db.Init()
	db.InitESDB()
	r.Use(ContextInjector("null", config))

	v1 := r.Group("/v1")
	{
		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/user/signin", user.Signin)
		v1.POST("/user/signup", user.Signup)
		v1.POST("/user/signout", user.Signout)

		/*** START ElasticSearch ***/
		es := new(controllers.EsController)
		v1.POST("/table", es.AddTable)
		v1.POST("/_search", es.Search)
		v1.POST("/_msearch", es.MSearch)

		/*** START Table Management ***/
		table := new(controllers.TableController)
		v1.POST("/mgmt/table", table.AddTable)
		v1.PUT("/mgmt/table", table.UpdateTable)
		v1.DELETE("/mgmt/table", table.DeleteTable)
	}

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/public/dist/index.html")
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run(":9000")
}
