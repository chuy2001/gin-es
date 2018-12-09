package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/chuy2001/gin-es/forms"
	"github.com/chuy2001/gin-es/models"

	"log"
	"strconv"
)

type InstanceController struct{}
var InstanceModel = new(models.InstanceModel)

//GetTable ...
func (ctrl InstanceController) GetInstance(c *gin.Context) {
	page_size := c.DefaultQuery("size", "10")
	page := c.DefaultQuery("page", "0")
	var instanceForm forms.InstanceForm

	if c.BindJSON(&instanceForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form", "form": instanceForm})
		c.Abort()
		return
	}

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

	instance, err := InstanceModel.GetInstance(size,from,instanceForm)
	if err == nil {
		c.JSON(200, gin.H{"message": "table get success","data":instance})
	} else {
		c.JSON(406, gin.H{"message": "table get failed", "error": err.Error()})
	}
	log.Println("TableController GetTable OK")
}

//AddInstance ...
// func (ctrl InstanceController) AddInstance(c *gin.Context) {
// 	var tableForm forms.TableForm

// 	if c.BindJSON(&tableForm) != nil {
// 		c.JSON(406, gin.H{"message": "Invalid form", "form": tableForm})
// 		c.Abort()
// 		return
// 	}

// 	tableForm.Creation_Time = time.Now().Format(time.RFC3339)
// 	tableForm.Creator_UserName = "admin"

// 	onetable, err := InstanceModel.AddInstance(tableForm)
// 	if err == nil {
// 		c.JSON(200, gin.H{"message": "table add success","data":onetable})
// 	} else {
// 		c.JSON(406, gin.H{"message": "table add failed", "error": err.Error()})
// 	}
// 	log.Println("TableController AddTable OK")
// }

//UpdateInstance ...
// func (ctrl InstanceController) UpdateInstance(c *gin.Context) {
// 	var tableForm forms.TableForm
// 	var id = c.Param("id")

// 	if c.BindJSON(&tableForm) != nil {
// 		c.JSON(406, gin.H{"message": "Invalid form", "form": tableForm})
// 		c.Abort()
// 		return
// 	}

// 	onetable, err := InstanceModel.UpdateInstance(id,tableForm)
// 	if err == nil {
// 		c.JSON(200, gin.H{"message": "table add success","data":onetable})
// 	} else {
// 		c.JSON(406, gin.H{"message": "table add failed", "error": err.Error()})
// 	}
// 	log.Println("TableController AddTable OK")
// }

//DeleteInstance ...
// func (ctrl InstanceController) DeleteInstance(c *gin.Context) {
// 	var id = c.Param("id")

// 	_,err := InstanceModel.DeleteInstance(id)
// 	if err == nil {
// 		c.JSON(200, gin.H{"message": "table delete success"})
// 	} else {
// 		c.JSON(406, gin.H{"message": "table delete failed", "error": err.Error()})
// 	}
// 	log.Println("TableController delete OK")
// }