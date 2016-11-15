package model
import
(
	"log"
	"github.com/jmoiron/sqlx"
)
//import "github.com/gin-gonic/gin"
//import "log"

type Buy struct {
	Id              int64
	DocNo           string
	DocDate         string
	VendorId        int
	SumOfItemAmount float32
	BeforeTaxAmount float32
	TaxAmount       float32
	TotalAmount     float32
}

type row struct {
	Col string
	Item string
}
type BuyTrans struct {
	Doc  Buy
	Item []StockCard
}

func (b *BuyTrans)New() (Response) {
	log.Println("Model Purchase New Start")
	r := Response{}
	r.Code = 200
	r.Message = "SUCCESS"
	dbconn := Connectdb()

	sql := `insert into buy (docno,docdate,vendorid,sumofitemamount,beforetaxamount,taxamount,totalamount) values(?,?,?,?,?,?,?)`
	x, err := dbconn.Exec(sql,
		b.Doc.DocNo,
		b.Doc.DocDate,
		b.Doc.VendorId,
		b.Doc.SumOfItemAmount,
		b.Doc.BeforeTaxAmount,
		b.Doc.TaxAmount,
		b.Doc.TotalAmount,
	)

	if err != nil {
		println("Exec err:", err.Error())
	} else {
		Id, _ := x.LastInsertId()
		log.Println("Insert Header of buy success")
		// Update transaction.Id from last insert
		b.Doc.Id = Id
		// Todo: Insert Detail of Document

		b.NewDetail(dbconn)
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
		return r
	}
//	return
}


func (b *BuyTrans)NewDetail( dbconn *sqlx.DB) {
	// PURCHASE IS Doctype :1  Get data from init.go  constants
	docType := PURCHASE

	log.Println("Start Buytrans.NewDetail insert ")
	log.Println(b.Doc)
	for k, _ := range b.Item {
		sql := `insert into stockcard (doctype,docid,docdate,qty,price,amount,locationid,unitid,itemid)
				values (?,?,?,?,?,?,?,?,?	)`

		b.Item[k].Amount = b.Item[k].Qty * b.Item[k].Price
		log.Println(sql)
		_, err := dbconn.Exec(sql,
			docType,
			b.Doc.Id,
			b.Doc.DocDate,
			b.Item[k].Qty,
			b.Item[k].Price,
			b.Item[k].Amount,
			b.Item[k].LocationId,
			b.Item[k].UnitId,
			b.Item[k].ItemId,
		)
		if err != nil {
			println("Exec err:", err.Error())
		}
		// call update stock
		stc := Stock{}
		stc.UpdateStock(b.Item[k].ItemId, b.Item[k].LocationId, b.Item[k].Amount, PURCHASE, b.Item[k].Qty, dbconn)

	}
}


//func (b *BuyTrans)DeleteBuy()(Response){
//	log.Println("Model Purchase Delete Start")
//	r := Response{}
//	r.Code = 200
//	r.Message = "SUCCESS"
//
//	dbconn := Connectdb()
//
//	// Cancel Header of Buy trans
//	sql := `Update buy set iscancel=1 where docid = ?`
//	_, err := dbconn.Exec(sql,b.Doc.Id)
//	if err != nil {
//		println("Exec err:", err.Error())
//		r.Message=(err)
//
//	}
//
//
//	// Cancel Detail of Buy item Loop for each item - Decrease Stock (-)
//	stkcard := []StockCard{}
//	sql = `select * from stockcard where docid =?`
//	 err = dbconn.Select(&stkcard,sql,b.Doc.Id)
//
//	if err != nil {
//		println("Exec err:", err.Error())
//		r.Message=(err)
//
//	}
//
//	// loop for cancel & Update stock
//
//
//
//	return r
//}