package controller

import (
		m "github.com/mrtomyum/inventory/model"
		"github.com/gin-gonic/gin"
		//"github.com/jmoiron/sqlx"
	"regexp"
	"strconv"
	"log"
)
var leadingInt = regexp.MustCompile(`^[-+]?\d+`)
var i []m.ItemProfile



func GetAllItem(c *gin.Context){


	//c.JSON(200, gin.H{"Message":"SUCCESS"})
	log.Println("controller get-All-Item start")
	item := m.Item{}
	i = item.GetAllItem()
	log.Println(i)
	c.JSON(200, i)

}


func GetItem(c *gin.Context){
	search_id := c.Params.ByName("id")

	//c.JSON(200, gin.H{"Message":"SUCCESS"})
	log.Println("controller Getitem start")
	x := c.Params.ByName("id")

	item := m.Item{}
	//y :=	strconv.ParseInt(x, 10, 64)

	log.Printf("Search Item by id : %v",search_id)
	// convert to int parameter to Y variable
	y, _ := strconv.Atoi(x)
	i := item.GetItem(y)
	c.JSON(200, i)

}


func  InsertItem(c *gin.Context){
	log.Println("controller insert item start")
	i := m.Item{}
	res := i.NewItem(c)
	c.JSON(200, res)
}

func UpdateItem(c *gin.Context){
	log.Println("Controller Update Item Start")
	i := m.Item{}
	res := i.UpdateItem(c)
	c.JSON(200,res)
}