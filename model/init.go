package model

// next we define the constants of type Day
// we use iota to provide int values to each of the contant types
type Doctype int
const (
	BIGINBAL Doctype = 0 + iota
	PURCHASE
	SALE
	ADJUST
	TRANSFERIN
	TRANSFEROUT

)