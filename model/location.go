package model
import (
	"log"
)
type Location struct {
	Id int
	Name string
	Address string
	Type int
}


func (l *Location)GetAllLocation()([]Location) {
	lc := []Location{}
	db := Connectdb()
	sql := `SELECT * FROM location`
	log.Println("begin Model.GetAllLocation")
	db.Select(&lc, sql)
	log.Println(lc)
	return lc

}