package main
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
//	m "github.com/mrtomyum/inventory/model"
	x "github.com/mrtomyum/inventory/controller"
//	"log"
//	"net/http"
//	"golang.org/x/crypto/bcrypt"
//	"github.com/denisenkom/go-mssqldb"
	//"log"
)





func main() {


	r := gin.New()

	//r.GET("/", Hello)
	r.GET("/items", x.GetItem)
	//r.POST("/item", InsertItem)

	// Master data part
	r.GET("/units", x.GetUnit)
	//r.POST("/users/login", UserLogin)
	r.Run(":9000")
}



//func InsertItem(c *gin.Context){
//	newitem := m.Item{}
//	c.BindJSON(&newitem)
//	res := m.Response{}
//	sql := `insert into item (code,name,unitid,price) values(?,?,?,?)`
//	_,err := DB.Exec(sql,
//			newitem.Code,
//			newitem.Name,
//			newitem.Unitid,
//			newitem.Price)
//	log.Println(err)
//	res.Code = 200
//	if err != nil {
//			res.Code = 200 //not modified
//			c.String(res.Code, err.Error())
//		return
//	} else {
//		c.JSON(res.Code, gin.H{"status":"SUCCESS"})
//		return
//	}
//
//
//	return
//}

