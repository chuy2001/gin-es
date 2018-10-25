package db

import (
	"fmt"
	"log"

	"github.com/chuy2001/gorqlite"
)

const (
	//DbUser ...
	DbUser = "admin"
	//DbPassword ...
	DbPassword = "admin"
)

var db *gorqlite.Connection

//Init ...
func Init() {

	dbinfo := fmt.Sprintf("http://%s:%s@localhost:4001/", DbUser, DbPassword)
	fmt.Printf("sqlite info %s\n", dbinfo)

	var err error
	conn, err := gorqlite.Open(dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	db = &conn

}

//GetDB ...
func GetDB() *gorqlite.Connection {
	return db
}
