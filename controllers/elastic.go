package controllers

import (
	"github.com/gin-gonic/gin"
	. "gopkg.in/olivere/elastic.v5"
	"github.com/chuy2001/gin-es/db"

	"context"
	"log"
	"encoding/json"
	"net/url"
	"io/ioutil"
)

type EsController struct{}


type Product struct {
    Name      string  `json:"name"`
    ProductID int64   `json:"product_id,string"`
    Number    int     `json:"number,string"`
    Price     float64 `json:"price,string"`
    IsOnSale  bool    `json:"is_on_sale,string"`
}

//AddTable ...
func (ctrl EsController) AddTable(c *gin.Context) {
	// client := c.MustGet("ESClient").(*es.Client)
	client := db.GetESDB()
	q2 := NewTermQuery("tags", "golang")
	sreq2 := NewSearchRequest().Index("elastic-test").Type("tweet").
		Source(NewSearchSource().Query(q2))

	searchResult, err := client.MultiSearch().
		Add(sreq2).
		Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if searchResult.Responses == nil {
		log.Fatal("expected responses != nil; got nil")
	}
	// if len(searchResult.Responses) != 2 {
	// 	log.Fatalf("expected 2 responses; got %d", len(searchResult.Responses))
	// }
	c.JSON(200, gin.H{"Message": searchResult.Responses})
}

//search ...
func (ctrl EsController) Search(c *gin.Context) {
	client := db.GetESDB()

	q2 := NewTermQuery("tags", "golang")

	sreq2 := NewSearchRequest().Index("elastic-test").Type("tweet").
		Source(NewSearchSource().Query(q2))

	searchResult, err := client.MultiSearch().
		Add(sreq2).
		Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if searchResult.Responses == nil {
		log.Fatal("expected responses != nil; got nil")
	}
	// if len(searchResult.Responses) != 2 {
	// 	log.Fatalf("expected 2 responses; got %d", len(searchResult.Responses))
	// }
	c.JSON(200, gin.H{"Message": searchResult.Responses})
}

//msearch ...
func (ctrl EsController) MSearch(c *gin.Context) {
	es := db.GetESDB()

	// 方法一，格式化msearch语法
	// q1 := `{"size":1,"query":{"bool":{"filter":[{"range":{"@timestamp":{"gte":"1545575756885","lte":"1545576056885","format":"epoch_millis"}}},{"query_string":{"analyze_wildcard":true,"query":"*"}}]}},"aggregations":{"2":{"date_histogram":{"interval":"10s","field":"@timestamp","min_doc_count":0,"extended_bounds":{"min":"1545575756885","max":"1545576056885"},"format":"epoch_millis"},"aggregations":{"1":{"avg":{"field":"mem.available_percent"}}}}}}`
	// p1 := &SearchSource{}
	// err1 := json.Unmarshal([]byte(q1), p1)
	// log.Println(err1)

	// 方法二  按ES语法搜索
	// q3 := RawStringQuery(`{"match_all" : {}}`)
	// sreq2 := NewSearchRequest().Index("telegraf*").
	// 	Source(NewSearchSource().Query(q3))
	// searchResult1, err1 := es.MultiSearch().
	// 	Add(sreq2).
	// 	Do(context.TODO())
	// if err1 != nil {
	// 	log.Fatal(err1)
	// }
	// if searchResult1.Responses == nil {
	// 	log.Println("expected responses != nil; got nil")
	// }

	// 方法三  直接转发前端的msearch请求
	// Build url
	path := "/_msearch"

	// Parameters
	params := make(url.Values)
	// if s.pretty {
	// 	params.Set("pretty", fmt.Sprintf("%v", s.pretty))
	// }

	body,_ := ioutil.ReadAll(c.Request.Body)
	log.Println("---body/--- \r\n "+string(body))
	
	// Get response
	res, err := es.PerformRequest(c, "GET", path, params,string(body))
	if err != nil {
		log.Println(err)
	}
	searchResult := new(MultiSearchResult)
	if err := json.Unmarshal(res.Body, searchResult); err != nil {
		log.Println(err)
	}

	if searchResult.Responses == nil {
		log.Println("expected responses != nil; got nil")
	}
	c.JSON(200, gin.H{"Message": searchResult.Responses})
}


// func (ctrl EsController) MSearch(c *gin.Context) {
// 	client := db.GetESDB()

// 	q2 := es.NewTermQuery("host", "DESKTOP-GGBB3LD")

// 	sreq2 := es.NewSearchRequest().Index("telegraf*").Type("metrics").
// 		Source(es.NewSearchSource().Query(q2))

// 	searchResult, err := client.MultiSearch().
// 		Add(sreq2).
// 		Do(context.TODO())
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if searchResult.Responses == nil {
// 		log.Fatal("expected responses != nil; got nil")
// 	}
// 	// if len(searchResult.Responses) != 2 {
// 	//  log.Fatalf("expected 2 responses; got %d", len(searchResult.Responses))
// 	// }
// 	c.JSON(200, gin.H{"Message": searchResult.Responses})
// }