package model
import (
	//"github.com/gin-gonic/gin"
	"log"
	//"github.com/jmoiron/sqlx"

	//"github.com/jmoiron/sqlx"
//	"github.com/jmoiron/sqlx"
)

type Item struct {
	Code string
	Name string
	Unitid int
	Price float32
}


type ItemProfile struct {
	Id       int    `json:"id" db:"id"`
	Code     string `json:"code" `
	Name     string `json:"name"`
	Stockqty float32 `json:"stockqty"`
	Amount   float32 `json:"amount" `
	Unitname string `json:unitname `
}


func GetAllItem() ([]ItemProfile){

	items :=  []ItemProfile{}
	dbconn := Connectdb()
	log.Println("begin model.Getallitem")
	sql := `SELECT * FROM itemProfile  `
	dbconn.Select(&items, sql)
	log.Println(items)
	//c.JSON(200, items)
	return items
}

