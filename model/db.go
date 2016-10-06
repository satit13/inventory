package model
import (
	"github.com/jmoiron/sqlx"
	"log"
)
const (
	DB_HOST = "tcp(nava.work:3306)"
	DB_NAME = "item"
	DB_USER = "root"
	DB_PASS = "mypass"
)

var DB *sqlx.DB


func Connectdb()(*sqlx.DB) {
	var dsn = DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?parseTime=true"
	DB = sqlx.MustConnect("mysql", dsn)

	log.Println(DB.Ping())
	return DB
}