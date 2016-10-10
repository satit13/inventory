package controller
import (

	m "github.com/mrtomyum/inventory/model"
	"github.com/gin-gonic/gin"
	//"log"
)



func GetUnit(c *gin.Context){
	u := m.Unit{}
	units := u.GetAllUnit()
	c.JSON(200, units)
	return
}
