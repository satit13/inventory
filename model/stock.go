package model
import (
	"log"
	"github.com/jmoiron/sqlx"
)
type Stock struct {
	Id int64
	Itemid int64
	LocationId int64
	Qty float32
	Amount float32
	Averagecost float32


}

func(i *Stock)UpdateStock(itemid int64,locid int64,amount float32,doctype Doctype,qty float32,dbconn *sqlx.DB){
	log.Println("< ------ Begin Stockcard.Update process")

	//check existing record
	var sql string
		sql = `select *  from stock where itemid=? and locationid = ?`
		log.Println(sql)
		st := []Stock{}
		dbconn.Select(&st, sql,	itemid,	locid,)

		log.Printf("Check Existing Data : %v Record",len(st))
	if len(st) <= 0 {

		sql = `insert into stock (itemid,locationid) values (?,?)`
		log.Println(sql)
		_,err := dbconn.Exec(sql,itemid,locid,)
		if err != nil {
			println("Exec err:", err.Error())
		}
	}

		sql = `update stock set qty = qty+?,amount = amount+?  where itemid=? and locationid = ?`

	log.Println(sql)
	_,err := dbconn.Exec(sql,
		qty,
		amount,
		itemid,
		locid,
	)

	if err != nil {
		println("Exec err:", err.Error())
	}


}
