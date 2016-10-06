package main
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	x "github.com/mrtomyum/inventory/controller"
)


func main() {


	r := gin.New()
	//r.GET("/", Hello)
	r.GET("/items", x.GetItem)
	r.POST("/item", x.InsertItem)
	r.PUT("/item", x.UpdateItem)

	// Buy transaction
	// todo: is coding ... Not completed ...

	r.POST("/buy" , x.NewBuy)

	// Environment Master data part
	r.GET("/units", x.GetUnit)
	r.GET("/locations", x.GetLocation)
	//r.POST("/users/login", UserLogin)

	r.Run(":9000")


}




