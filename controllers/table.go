package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/chuy2001/gin-es/forms"
	"github.com/chuy2001/gin-es/models"

	"log"
	"strconv"
	"time"
)

type TableController struct{}
var TableModel = new(models.TableModel)

//GetTable ...
func (ctrl TableController) GetTable(c *gin.Context) {
	page_size := c.DefaultQuery("size", "10")
	page := c.DefaultQuery("page", "0")
	search := c.Query("search") 

	//字符串转换成整数
	size,err := strconv.Atoi(page_size)
	if err != nil{
		log.Println("字符串转换成整数失败")
		c.JSON(406, gin.H{"message": "table add failed", "error": err.Error()})
	}
	from,err := strconv.Atoi(page)
	if err != nil{
		log.Println("字符串转换成整数失败")
		c.JSON(406, gin.H{"message": "table add failed", "error": err.Error()})
	}

	table, err := TableModel.GetTable(size,from,search)
	if err == nil {
		c.JSON(200, gin.H{"message": "table get success","data":table})
	} else {
		c.JSON(406, gin.H{"message": "table get failed", "error": err.Error()})
	}
	log.Println("TableController GetTable OK")
}

//AddTable ...
func (ctrl TableController) AddTable(c *gin.Context) {
	var tableForm forms.TableForm

	if c.BindJSON(&tableForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": tableForm})
		c.Abort()
		return
	}

	tableForm.Creation_Time = time.Now().Format(time.RFC3339)
	tableForm.Creator_UserName = "admin"

	onetable, err := TableModel.AddTable(tableForm)
	if err == nil {
		c.JSON(200, gin.H{"message": "table add success","data":onetable})
	} else {
		c.JSON(406, gin.H{"message": "table add failed", "error": err.Error()})
	}
	log.Println("TableController AddTable OK")
}

//UpdateTable ...
func (ctrl TableController) UpdateTable(c *gin.Context) {
	var tableForm forms.TableForm
	var id = c.Param("id")

	if c.BindJSON(&tableForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": tableForm})
		c.Abort()
		return
	}

	onetable, err := TableModel.UpdateTable(id,tableForm)
	if err == nil {
		c.JSON(200, gin.H{"message": "table add success","data":onetable})
	} else {
		c.JSON(406, gin.H{"message": "table add failed", "error": err.Error()})
	}
	log.Println("TableController AddTable OK")
}

//DeleteTable ...
func (ctrl TableController) DeleteTable(c *gin.Context) {
	var id = c.Param("id")

	_,err := TableModel.DeleteTable(id)
	if err == nil {
		c.JSON(200, gin.H{"message": "table delete success"})
	} else {
		c.JSON(406, gin.H{"message": "table delete failed", "error": err.Error()})
	}
	log.Println("TableController delete OK")
}