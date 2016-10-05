package controller
import (

	m "github.com/mrtomyum/inventory/model"
	"github.com/gin-gonic/gin"
	//"log"
)



func GetUnit(c *gin.Context){

	units := m.GetAllUnit()


	c.JSON(200, units)
	return
}
