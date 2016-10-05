package model
type Unit struct {
	Id int
	Name string
}


func GetAllUnit()([]Unit){

	dbconn := Connectdb()
	units := []Unit{}
	sql := `SELECT * FROM unit  `
	dbconn.Select(&units, sql)
	return units
}