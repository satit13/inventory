package main
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
//	m "github.com/mrtomyum/inventory/model"
	x "github.com/mrtomyum/inventory/controller"
)


func main() {
	//m.TestConnectMsSql()


	r := gin.New()
	//r.GET("/", Hello)
	r.GET("/items", x.GetAllItem)
	r.GET("/item/:id" ,x.GetItem)

	r.POST("/item", x.InsertItem)
	r.PUT("/item", x.UpdateItem)


	// Buy transaction
	// todo: is coding ... Not completed ...

	r.POST("/buy", x.NewBuy)
	r.DELETE("/buy", x.DeleteBuy)

	// Environment Master data part
	r.GET("/units", x.GetUnit)
	r.GET("/locations", x.GetLocation)
	r.GET("/stocks",x.GetAllStock)
	//r.POST("/users/login", UserLogin)


	r.Run(":9000")


}




