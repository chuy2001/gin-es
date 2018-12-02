package models

import (
	"fmt"
	"log"
	"context"
	"encoding/json"

	"github.com/chuy2001/gin-es/db"
	"github.com/chuy2001/gin-es/forms"
	. "gopkg.in/olivere/elastic.v5"
)
var _DefaultIndex = "default_class_index"
var _DefaultType  = "default_class_type"
var Mapping = `{
	"settings":{
		"number_of_shards":5,
		"number_of_replicas":1
	},
	"mappings":{
		"default_class_type":{
			"properties":{
				"name":{"type":"keyword"},
				"alias":{"type":"string"},
				"readme":{"type":"string"},
				"creation_time":{"type":"date"},
				"creator_username":{"type":"string"},
				"creator":{"type":"string"}
			}
		}
	}
}`

//TableModel ...
type TableModel struct{}

//GetTable ...
func (m TableModel) GetTable(page_size int,page int,search string) (table []forms.TableFormWithID, err error) {
	es := db.GetESDB()
	m.init(_DefaultIndex,Mapping)

	all := NewMatchAllQuery()
	searchResult, err := es.Search().Index(_DefaultIndex).Query(all).Do(context.TODO())
	if err != nil {
		log.Println("Get Table failed")
		return table, err
	}
	if searchResult.Hits == nil {
		log.Printf("expected SearchResult.Hits != nil; got nil")
	}
	var lTableForm = make([]forms.TableFormWithID,0)

	var vTableForm forms.TableFormWithID
	for _, hit := range searchResult.Hits.Hits {
		if hit.Index != _DefaultIndex {
			log.Printf("expected SearchResult.Hits.Hit.Index = %q; got %q", _DefaultIndex, hit.Index)
		}
		err := json.Unmarshal(*hit.Source, &vTableForm)
		if err != nil {
			log.Println("Get Source failed")
		}
		vTableForm.ID = hit.Id

		lTableForm = append(lTableForm, vTableForm)
	}

	return lTableForm, nil
}

//AddTable ...
func (m TableModel) AddTable(form forms.TableForm) (table forms.TableFormWithID, err error) {
	log.Println("Add Table index")
	es := db.GetESDB()
	m.init(_DefaultIndex,Mapping)

	// Add doc
	put1, err := es.Index().Index(_DefaultIndex).Type(_DefaultType).BodyJson(form).Do(context.Background())
    if err != nil {
		log.Println("Add Table failed")
	}
	_, err = es.Flush().Index(_DefaultIndex).Do(context.TODO())
	if err != nil {
		log.Println("Add Table Flush failed")

	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)
	OneTable, err:=m.DocQuery(_DefaultIndex,_DefaultType,put1.Id)
	
	return  OneTable, nil
}

//UpdateTable ...
func (m TableModel) UpdateTable(id string, form forms.TableForm) (table forms.TableFormWithID, err error) {
	log.Println("Add Table index")
	es := db.GetESDB()
	m.init(_DefaultIndex,Mapping)

	// update doc
	put1, err := es.Index().Index(_DefaultIndex).Type(_DefaultType).Id(id).BodyJson(form).Do(context.Background())
    if err != nil {
		log.Println("Add Table failed")
	}
	_, err = es.Flush().Index(_DefaultIndex).Do(context.TODO())
	if err != nil {
		log.Println("Add Table Flush failed")
	}
	fmt.Printf("Indexed tweet %s to index s%s, type %s\n", put1.Id, put1.Index, put1.Type)
	OneTable, err:=m.DocQuery(_DefaultIndex,_DefaultType,put1.Id)
	
	return  OneTable, nil
}

//DeleteTable ...
func (m TableModel) DeleteTable(id string) (ret string, err error) {
	log.Println("DeleteTable start")	
	es := db.GetESDB()

	// Delete doc
	res,err:=es.Delete().Index(_DefaultIndex).Type(_DefaultType).Id(id).Do(context.Background())
	if err != nil{
		log.Println("Del Table failed")
		return "ok",err
	}
	if res.Found != true {
		log.Printf("expected Found = true; got %v", res.Found)
	}
	_, err = es.Flush().Index(_DefaultIndex).Do(context.TODO())
	if err != nil {
		log.Println("Del Table Flush failed")
		return "ok",err
	}
	ret = "DeleteTable success"
	return ret, err
}


//DeleteAttrTable ...
func (m TableModel) DeleteAttrTable(form forms.TableForm) ( err error) {
	es := db.GetESDB()
	fmt.Println("DeleteTable Success",es.GetMapping())

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
	} else {
		fmt.Println("Index is not found!")
	}
	return nil
}

func (m TableModel) init(name string, mapping string) (err error) {
	es := db.GetESDB()	
	// Exists
	indexExists, err := es.IndexExists(_DefaultIndex).Do(context.TODO())
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

//QueryDoc ...
func (m TableModel) DocQuery(_index string, _type string, _id string) (table forms.TableFormWithID, err error) {
	log.Println("Add Table index")
	es := db.GetESDB()

	// Query doc
	res,err:=es.Get().Index(_index).Type(_type).Id(_id).Do(context.Background())
	if err != nil {
		log.Println(err)
	}
	if res.Found != true {
		log.Printf("expected Found = true; got %v", res.Found)
	}

	var vTableForm forms.TableFormWithID
	err = json.Unmarshal(*res.Source, &vTableForm)
	if err != nil {
		log.Println("Get Source failed")
	}
	vTableForm.ID = res.Id

	return vTableForm, nil
}