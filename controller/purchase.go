package controller
import (
	m "github.com/mrtomyum/inventory/model"
	"github.com/gin-gonic/gin"
//"github.com/jmoiron/sqlx"

	"log"
)

func NewBuy(c *gin.Context){
	log.Println("Begin Controller  POST New Buy Transaction")
	b := m.Buy{}
	c.BindJSON(&b)
	log.Println(b)
	result := m.Response{}

	result = b.New(b)
	log.Println(result)
	c.JSON(200, result)
}