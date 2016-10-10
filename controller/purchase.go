package controller
import (
	m "github.com/mrtomyum/inventory/model"
	"github.com/gin-gonic/gin"
//"github.com/jmoiron/sqlx"

	"log"
)

func NewBuy(c *gin.Context) {
	log.Println("Begin Controller  POST New Buy Transaction")
	b := m.BuyTrans{}

	c.BindJSON(&b)
	//	log.Println(b.Doc.DocNo)
	//	log.Println(b.Doc.DocDate)
	//	log.Println(b.Doc.TotalAmount)
	//	log.Println(b.Doc.VendorId)
	//	log.Println("--item detail--")
	//	log.Println(b.Item[0].ItemId)
	//	log.Println(b.Item[0].Price)

	log.Println(b.Item)


	result := m.Response{}

	// call model.purchase.newbuy

	result = b.New()
	log.Println(result)
	c.JSON(200, result)
}


func DeleteBuy(c *gin.Context) {
	log.Println("Begin Controller  DELETE  Buy Transaction")
	b := m.BuyTrans{}

	c.BindJSON(&b)
	//	log.Println(b.Doc.DocNo)
	//	log.Println(b.Doc.DocDate)
	//	log.Println(b.Doc.TotalAmount)
	//	log.Println(b.Doc.VendorId)
	//	log.Println("--item detail--")
	//	log.Println(b.Item[0].ItemId)
	//	log.Println(b.Item[0].Price)

	log.Println(b.Item)
	result := m.Response{}

	// call model.purchase.newbuy

//	result = b.DeleteBuy()
	log.Println(result)
	c.JSON(200, result)
}