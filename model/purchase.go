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

func(b *Buy)New(new Buy )(Response){
	log.Println("Model Purchase New Start")
	r := Response{}
	r.Code = 200
	r.Message = "SUCCESS"
	return r
}



