package models

import (
	"log"
	"context"
	"encoding/json"

	"github.com/chuy2001/gin-es/db"
	"github.com/chuy2001/gin-es/forms"
	. "gopkg.in/olivere/elastic.v5"
)
var _InstanceIndex = "default_instance_index"
var _InstanceType  = "default_instance_type"
var _InstanceMapping = `{
	"settings":{
		"number_of_shards":5,
		"number_of_replicas":1
	},
	"mappings":{
		"default_instance_type":{
			"properties":{
				"name":{"type":"keyword"},
				"id":{"type":"keyword"}
			}
		}
	}
}`

//InstanceModel ...
type InstanceModel struct{}

//GetTable ...
func (m InstanceModel) GetInstance(page_size int,page int,from forms.InstanceForm) (table forms.InstanceRsp, err error) {
	es := db.GetESDB()
	if from.Type != "" {
		_InstanceIndex = from.Type
	}
	m.init(_InstanceIndex,_InstanceMapping)

	all := NewMatchAllQuery()
	searchResult, err := es.Search().Index(_InstanceIndex).Query(all).Do(context.TODO())
	if err != nil {
		log.Println("Get Table failed",err)
		return table, err
	}
	if searchResult.Hits == nil {
		log.Printf("expected SearchResult.Hits != nil; got nil")
	}

	var vTableForm forms.InstanceRsp
	vTableForm.Page.Total = searchResult.Hits.TotalHits
	for _, hit := range searchResult.Hits.Hits {
		if want, got := "", hit.Index; want != got {
			log.Printf("expected index %q, got %q", want, got)
		}
		item := make(map[string]interface{})
		err := json.Unmarshal(*hit.Source, &item)
		if err != nil {
			log.Println(err)
		}
		vTableForm.List = append(vTableForm.List, item)
	}
	return vTableForm, nil
}


func (m InstanceModel) init(name string, mapping string) (err error) {
	es := db.GetESDB()	
	// Exists
	indexExists, err := es.IndexExists(name).Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if !indexExists {
		// Create index
		createIndex, err := es.CreateIndex(name).Body(mapping).Do(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		if createIndex == nil {
			log.Printf("createIndex fail response; got: %v", createIndex)
		}
		if !createIndex.Acknowledged {
			log.Println("createIndex Acknowledged")
		}
	}
	return err
}
