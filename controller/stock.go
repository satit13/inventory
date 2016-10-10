package controller
import (

	m "github.com/mrtomyum/inventory/model"
	"github.com/gin-gonic/gin"
//"log"
	"log"
)



func GetAllStock(c *gin.Context){
	log.Println("Begin Controller.stock.GetAllStock")
	s := m.Stockprofile{}

	i := s.GetAllStock()
	log.Println("result to controller ")
	log.Println(i)
	c.JSON(200, i)
	return
}