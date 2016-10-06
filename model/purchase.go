package model

type Buy struct {
	DocNo string
	DocDate string
	DueDate string
	VendorId string
	SumOfItemAmount float32
	DiscountAmount float32
	BeforeTaxAmount float32
	TaxAmount float32
	TotalAmount float32
	TaxType int
	IsCancel int
	TaxRate float32
}

