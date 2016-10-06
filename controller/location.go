package controller
import (

	m "github.com/mrtomyum/inventory/model"
	"github.com/gin-gonic/gin"
	//"log"

	"log"
)


func GetLocation(c*gin.Context){
	l := []m.Location{}

	l = m.GetAllLocation()
	log.Println(l)
	c.JSON(200, l)
	return
}
