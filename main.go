package main
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
//	"log"
//	"net/http"
//	"golang.org/x/crypto/bcrypt"
//	"github.com/denisenkom/go-mssqldb"

	"log"
)
const (
	DB_HOST = "tcp(nava.work:3306)"
	DB_NAME = "item"
	DB_USER = "root"
	DB_PASS = "mypass"
)
type ItemProfile struct {
	Id       int    `json:"id" db:"id"`
	Code     string `json:"code" `
	Name     string `json:"name"`
	Stockqty float32 `json:"stockqty"`
	Amount   float32 `json:"amount" `
	Unitname string `json:unitname `
}


type Item struct {
	Code string
	Name string
	Unitid int
	Price float32
}

type Unit struct {
	Id int
	Name string
}


var db *sqlx.DB

func main() {
	var dsn = DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?parseTime=true"
	db = sqlx.MustConnect("mysql", dsn)

	r := gin.New()
	//r.GET("/", Hello)
	r.GET("/items", getItem)
	r.POST("/item", InsertItem)

	// Master data part
	r.GET("/units", getUnit)
	//r.POST("/users/login", UserLogin)
	r.Run(":9000")
}

func getItem(c *gin.Context){
	items := []ItemProfile{}
	sql := `SELECT * FROM itemProfile  `
	db.Select(&items, sql)
	log.Println()
	c.JSON(200, items)
	return
}


func getUnit(c *gin.Context){
	units := []Unit{}
	sql := `SELECT * FROM unit  `
	error := db.Select(&units, sql)
	log.Println(error)
	c.JSON(200, units)
	return
}


func InsertItem(c *gin.Context){
	newitem := Item{}
	c.BindJSON(&newitem)

	sql := `insert into item (code,name,unitid,price) values(?,?,?,?)`
	_,err := db.Exec(sql,
			newitem.Code,
			newitem.Name,
			newitem.Unitid,
			newitem.Price)
	log.Println(err)

	return
}

