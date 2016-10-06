package controller

import (
		m "github.com/mrtomyum/inventory/model"
		"github.com/gin-gonic/gin"
		//"github.com/jmoiron/sqlx"

	"log"
)

var i []m.ItemProfile


func GetItem(c *gin.Context){

	//c.JSON(200, gin.H{"Message":"SUCCESS"})
	log.Println("controller getitem start")
	i = m.GetAllItem()
	log.Println(i)
	c.JSON(200, i)

}

func  InsertItem(c *gin.Context){
	log.Println("controller insert item start")
	res := m.NewItem(c)
	c.JSON(200, res)
}

func UpdateItem(c *gin.Context){
	log.Println("Controller Update Item Start")
	i := m.Item{}
	res := i.UpdateItem(c)
	c.JSON(200,res)
}