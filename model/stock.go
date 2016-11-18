package model
import (
	"log"
	"github.com/jmoiron/sqlx"
)
type Stock struct {
	Id          int64
	Itemid      int64
	LocationId  int64
	Qty         float32
	Amount      float32
	Averagecost float32
}

type Stockprofile struct {
	Id int64
	ItemId int64
	LocationId int64
	Qty float32
	Itemname string
	Locationname string

}


func (i *Stockprofile)GetAllStock()([]Stockprofile){

	log.Println("Begin Model.Stockprofile.GetAllStock")
	dbconn := Connectdb()
	rs := []Stockprofile{}

	sql := `select id,itemid,locationid,qty,itemname,locationname from StockProfile `
	log.Printf(sql)
	log.Println("---------------")
	error := dbconn.Select(&rs, sql)
	if error != nil {
		log.Println(error)
	}
	return rs
}

func (i *Stock)UpdateStock(doctype int,sc StockCard, dbconn *sqlx.DB) {
	log.Println("< ------ Begin Stockcard.Update process")

	//check existing record
	var sql string
	sql = `select *  from stock where itemid=? and locationid = ? `
	log.Println(sql)
	st := []Stock{}

	dbconn.Select(&st, sql, sc.ItemId, sc.LocationId )

	log.Printf("Check Existing Data : %v Record ", len(st))
	//log.Printf("Exists data id : %v",st[1].Id)
	if len(st) <= 0 {

		sql = `insert into stock (itemid,locationid) values (?,?)`
		log.Println(sql)
		_, err := dbconn.Exec(sql,sc.ItemId, sc.LocationId )
		if err != nil {
			println("Exec err:", err.Error())
		}
	}

	println("IsCancel : ",sc.Iscancel)
	println("Doctype  : ",sc.DocType)


	switch
	{
		case doctype == 1:
			println("Doctype : ",sc.DocType)
			if sc.Iscancel == 0	{
					sql = `update stock set qty = qty+?,amount = amount+?  where itemid=? and locationid = ?`
					log.Println(sql)
					_, err := dbconn.Exec(sql,
						sc.Qty,
						sc.Amount,
						sc.ItemId,
						sc.LocationId,	)
					if err != nil {
						println("Exec err2:", err.Error())
					}
				}
			// cancel
			if sc.Iscancel == 1 {
				sql = `update stock set qty = qty-?,amount = amount-?  where itemid=? and locationid = ?`
				log.Println("BEGIN IS CANCEL = 1 , PURCHASE")
				_, err := dbconn.Exec(sql,
					sc.Qty,
					sc.Amount,
					sc.ItemId,
					sc.LocationId,	)
					if err != nil {
						println("Exec err2:", err.Error())
					}
				}
		}




	// update stock in item table
	sql = `update item set qty = (select sum(qty) from stock where itemid = ?) ,
			amount = (select sum(amount) from stock where itemid = ?)	where id=?`
	log.Println(sql)
	_, err := dbconn.Exec(sql,
		sc.ItemId,
		sc.ItemId,
		sc.ItemId,
	)

	if err != nil {
		println("Exec err:", err.Error())
	}


}
