package controllers

import (
	"github.com/gin-gonic/gin"
	es "gopkg.in/olivere/elastic.v5"

	"context"
	"log"
)

type EsController struct{}

//msearch ...
func (ctrl EsController) Search(c *gin.Context) {
	client := c.MustGet("ESClient").(*es.Client)

	q2 := es.NewTermQuery("tags", "golang")

	sreq2 := es.NewSearchRequest().Index("elastic-test").Type("tweet").
		Source(es.NewSearchSource().Query(q2))

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
	client := c.MustGet("ESClient").(*es.Client)

	q2 := es.NewTermQuery("tags", "golang")

	sreq2 := es.NewSearchRequest().Index("elastic-test").Type("tweet").
		Source(es.NewSearchSource().Query(q2))

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
	//  log.Fatalf("expected 2 responses; got %d", len(searchResult.Responses))
	// }
	c.JSON(200, gin.H{"Message": searchResult.Responses})
}
