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