package db

import (
	"fmt"
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

func SearchTable(path string, sqlStatement string) []Item {
	db, err := sqlx.Connect("sqlite", path)
	if err != nil {
		log.Fatal("ERROR connecting to DB:", err)
	}

	fmt.Println("sqlStatement =", sqlStatement)

	itemList := []Item{}
	err = db.Select(&itemList, sqlStatement)

	if err != nil {
		log.Fatal("ERROR in SearchTable: ", err)
	}
	return itemList
}