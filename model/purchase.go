package model
import "log"
//import "github.com/gin-gonic/gin"
//import "log"

type Buy struct {
	Id int
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


	dbconn := Connectdb()


	sql := `insert into buy (docno,docdate,vendorid,sumofitemamount,beforetaxamount,taxamount,totalamount) values(?,?,?,?,?,?,?)`
	_,err := dbconn.Exec(sql,
		trx.Doc.DocNo,
		trx.Doc.DocDate,
		trx.Doc.VendorId,
		trx.Doc.SumOfItemAmount,
		trx.Doc.BeforeTaxAmount,
		trx.Doc.TaxAmount,
		trx.Doc.TotalAmount,
		)
	log.Println(sql)

//	Check slice of item detail before insert to database
	for k, _ := range trx.Item {
		log.Println(trx.Item[k].ItemId)
		log.Println(trx.Item[k].Qty)
		log.Println(trx.Item[k].Price)
		log.Println("--Next Item--")
	}

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



