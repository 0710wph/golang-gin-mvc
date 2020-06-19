package models

import (
	"database/sql"
	"gin/database"
	"log"
)
var db *sql.DB = database.DB

type Bd struct {
	Id        int
	Types  string
	Bdtypes  string
	Title string
	Bddate string
	Data_from string
	Areaname string
	Street string
}

func (model *Bd) QueryBdlist(sqlstrn string) (bds []Bd, err error) {
	//bds := make([]Bd, 0)
	bds = make([]Bd, 0)
	rows, err := db.Query(sqlstrn)
	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()
	for rows.Next() {
		var bd Bd
		rows.Scan(&bd.Id, &bd.Types, &bd.Bdtypes, &bd.Title, &bd.Bddate, &bd.Data_from, &bd.Areaname, &bd.Street)
		bds = append(bds, bd)
	}

	if err = rows.Err(); err != nil {
		log.Println(err.Error())
	}
	return
}
