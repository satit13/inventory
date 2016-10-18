package model
//import "log"

type StockCard struct {
	Id         int
	DocType    int64
	DocId      int64
	ItemId     int64
	Qty        float32
	LocationId int64
	Price      float32
	UnitId     float32
	NetAmount  float32
	TaxAmount  float32
	Amount     float32
	Iscancel   int
	Row 	  []row
}


