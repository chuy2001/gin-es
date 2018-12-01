package models

import (
	"fmt"
	"log"
	"context"


	"github.com/chuy2001/gin-es/db"
	"github.com/chuy2001/gin-es/forms"
	// es "gopkg.in/olivere/elastic.v5"
)



//TableModel ...
type TableModel struct{}

//AddTable ...
func (m TableModel) AddTable(form forms.TableForm) ( err error) {
	es := db.GetESDB()
	fmt.Println("AddTable Success",es.GetMapping())

	mapping := `{
		"settings":{
			"number_of_shards":5,
			"number_of_replicas":1
		},
		"mappings":{
			"host":{
				"properties":{
					"tags":{
						"type":"string"
					},
					"location":{
						"type":"geo_point"
					}
				}
			}
		}
	}`
	
	// Exists
	indexExists, err := es.IndexExists(form.Name).Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if indexExists {
		log.Printf("index exists=%v", indexExists)
		deleIndex,err:=es.DeleteIndex(form.Name).Do(context.TODO())
		if err != nil{
			log.Fatal(err)
		}

		if !deleIndex.Acknowledged {
			fmt.Println("!deleIndex.Acknowledged")
		} else {
			fmt.Println("deleIndex.Acknowledged")
		}

	}

	// Create index
	createIndex, err := es.CreateIndex(form.Name).Body(mapping).Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if createIndex == nil {
		log.Printf("createIndex fail response; got: %v", createIndex)
	}
	if !createIndex.Acknowledged {
		log.Println("createIndex Acknowledged")
	}
	return  nil
}
