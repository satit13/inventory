package model
import
(
	"log"
	"github.com/jmoiron/sqlx"
)
//import "github.com/gin-gonic/gin"
//import "log"

type Buy struct {
	Id int64
	DocNo string
	DocDate string
	VendorId int
	SumOfItemAmount float32
	BeforeTaxAmount float32
	TaxAmount float32
	TotalAmount float32
}

type BuyTrans struct {
	Doc Buy
	Item []StockCard
}

func(b *BuyTrans)New(trx BuyTrans )(Response){
	log.Println("Model Purchase New Start")
	r := Response{}
	r.Code = 200
	r.Message = "SUCCESS"


	t :=trx

	dbconn := Connectdb()


	sql := `insert into buy (docno,docdate,vendorid,sumofitemamount,beforetaxamount,taxamount,totalamount) values(?,?,?,?,?,?,?)`
	x,err := dbconn.Exec(sql,
		t.Doc.DocNo,
		t.Doc.DocDate,
		t.Doc.VendorId,
		t.Doc.SumOfItemAmount,
		t.Doc.BeforeTaxAmount,
		t.Doc.TaxAmount,
		t.Doc.TotalAmount,
		)

	if err != nil {
		println("Exec err:", err.Error())
	} else {
		Id, _ := x.LastInsertId()
		log.Println("Insert Header of buy success")
		// Update transaction.Id from last insert
		t.Doc.Id = Id
		// Todo: Insert Detail of Document

		t.NewDetail(t,dbconn)
	}

//	Check slice of item detail before insert to database



	//log.Println(err)
	r.Code = 200

	if err != nil {
		r.Message = err.Error()
		//c.String(res.Code, err.Error())
		return r
	} else {
		r.Message = "SUCCESS"
		return  r
	}


	return r
}


func(trx *BuyTrans)NewDetail(ts BuyTrans, dbconn *sqlx.DB){
	// PURCHASE IS Doctype :1  Get data from init.go  constants
	docType := PURCHASE

	log.Println("Start Buytrans.NewDetail insert ")
	log.Println(ts.Doc)
	for k, _ := range trx.Item {
		sql := `insert into stockcard (doctype,docid,docdate,qty,price,amount,locationid,unitid,itemid)
				values (?,?,?,?,?,?,?,?,?	)`
		log.Println(sql)
		_,err := dbconn.Exec(sql,
			docType,
			ts.Doc.Id,
			ts.Doc.DocDate,
			trx.Item[k].Qty,
			trx.Item[k].Price,
			trx.Item[k].Amount,
			trx.Item[k].LocationId,
			trx.Item[k].UnitId,
			trx.Item[k].ItemId,
		)
		if err != nil {
			println("Exec err:", err.Error())
		}

		// call update stock
		stc := Stock{}
		stc.UpdateStock(trx.Item[k].ItemId,trx.Item[k].LocationId,trx.Item[k].Amount,PURCHASE,trx.Item[k].Qty,dbconn)

	}
}
