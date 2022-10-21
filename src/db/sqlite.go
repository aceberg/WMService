package db

import (
	. "github.com/aceberg/WMService/models"
	"github.com/jmoiron/sqlx"
	"log"
	_ "modernc.org/sqlite"
)

func execAny(path string, sqlStatement string) {
	db, err := sqlx.Connect("sqlite", path)
	if err != nil {
		log.Fatal("ERROR connecting to DB:", err)
	}

	_, err = db.Exec(sqlStatement)
	if err != nil {
		log.Fatal("ERROR in execAny: ", err)
	}
}

func SelectAll(path string) []Item {
	db, err := sqlx.Connect("sqlite", path)
	if err != nil {
		log.Fatal("ERROR connecting to DB:", err)
	}

	itemList := []Item{}
	err = db.Select(&itemList, "SELECT * FROM rs_tickets ORDER BY DATE DESC")
	if err != nil {
		log.Fatal("ERROR in SelectAll: ", err)
	}
	return itemList
}

func SelectList(path string, table string) []string {
	db, err := sqlx.Connect("sqlite", path)
	if err != nil {
		log.Fatal("ERROR connecting to DB:", err)
	}

	list := []string{}
	err = db.Select(&list, "SELECT NAME FROM $1 ORDER BY NAME ASC", table)
	if err != nil {
		log.Fatal("ERROR in SelectList: ", err)
	}
	return list
}
