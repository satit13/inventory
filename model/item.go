package model
import (
	//"github.com/gin-gonic/gin"
	"log"
	//"github.com/jmoiron/sqlx"

	//"github.com/jmoiron/sqlx"
//	"github.com/jmoiron/sqlx"
	"github.com/gin-gonic/gin"
)

type Item struct {
	Id int
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
	Price float32 `json:price `
}


func (i *Item)GetAllItem() ([]ItemProfile){

	items :=  []ItemProfile{}
	dbconn := Connectdb()
	log.Println("begin model.Getallitem")
	sql := `SELECT * FROM itemProfile  `
	dbconn.Select(&items, sql)
	log.Println(items)
	//c.JSON(200, items)
	return items
}

func (i *Item)NewItem(c *gin.Context) (Response ){
	log.Println("Begin Model New Item ")
	newitem := Item{}
	c.BindJSON(&newitem)
	dbconn := Connectdb()
	res := Response{}


	sql := `insert into item (code,name,unitid,price) values(?,?,?,?)`
	_,err := dbconn.Exec(sql,
		newitem.Code,
		newitem.Name,
		newitem.Unitid,
		newitem.Price)
	//log.Println(err)
	res.Code = 200

	if err != nil {
		res.Message = err.Error()
		//c.String(res.Code, err.Error())
		return res
	} else {

		res.Message = "SUCCESS"
		return  res
	}

}


func (i *Item)UpdateItem(c *gin.Context)(Response){
	log.Println("Begin Model Update Item ")
	UpdateItem := Item{}
	c.BindJSON(&UpdateItem)
	dbconn := Connectdb()
	res := Response{}


	sql := `update item	set code = ?,name = ?,unitid=?,	price=?		where id = ?`
	log.Println(sql)
	_,err := dbconn.Exec(sql,UpdateItem.Code,UpdateItem.Name,UpdateItem.Unitid,	UpdateItem.Price,UpdateItem.Id)
	log.Println("Update Item ")
	log.Println(UpdateItem.Code)
	res.Code = 200

	if err != nil {
		res.Message = err.Error()
		//c.String(res.Code, err.Error())
		return res
	} else {

		res.Message = "SUCCESS"
		return  res
	}


}