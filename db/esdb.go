package db

import (
	"log"
	"context"

	es "gopkg.in/olivere/elastic.v5"
)

var esdb *es.Client
const esURL = "http://127.0.0.1:9200"

//Init ...
func InitESDB() {

	log.Printf("Connecting to ES on: %v", esURL)
	esdb, err := es.NewClient(es.SetURL(esURL), es.SetSniff(false))
	if err != nil {
		// Handle error
		log.Fatalln("ES Connecting failed",err)
		panic(err)
	}

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := esdb.Ping(esURL).Do(context.Background())
	if err != nil {
		// Handle error
		log.Fatalln("ES Ping failed",err)
		panic(err)
	}
	log.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut
	esversion, err := esdb.ElasticsearchVersion(esURL)
	if err != nil {
		// Handle error
		log.Fatalln("ES ElasticsearchVersion failed",err)
		panic(err)
	}
	log.Printf("Elasticsearch version %s\n", esversion)
}

//GetDB ...
func GetESDB() *es.Client {
	return esdb
}
