package model

type Buy struct {
	Id int
	DocNo string
	DocDate string
	VendorId string
	SumOfItemAmount float32
	BeforeTaxAmount float32
	TaxAmount float32
	TotalAmount float32
}

