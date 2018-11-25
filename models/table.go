package models

import (
	"fmt"

	"github.com/chuy2001/gin-es/db"
	"github.com/chuy2001/gin-es/forms"
)


//TableModel ...
type TableModel struct{}

//AddTable ...
func (m TableModel) AddTable(form forms.TableForm) ( err error) {
	es := db.GetESDB()
	fmt.Println("AddTable Success",es.GetMapping())

	return  nil
}
