package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/chuy2001/gin-es/forms"
	"github.com/chuy2001/gin-es/models"

	"log"
)

type TableController struct{}
var TableModel = new(models.TableModel)

//AddTable ...
func (ctrl TableController) AddTable(c *gin.Context) {
	var tableForm forms.TableForm

	if c.BindJSON(&tableForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": tableForm})
		c.Abort()
		return
	}

	err := TableModel.AddTable(tableForm)
	if err == nil {
		c.JSON(200, gin.H{"message": "table add success"})
	} else {
		c.JSON(406, gin.H{"message": "table add failed", "error": err.Error()})
	}
	log.Println("TableController AddTable OK")
}
