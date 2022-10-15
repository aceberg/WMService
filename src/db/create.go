package db

import (
	"log"
	"os"
)

func CreateDB(path string) {
	if _, err := os.Stat(path); err == nil {
		log.Println("INFO: DB exists")
	} else {
		sqlStatement := `CREATE TABLE IF NOT EXISTS rs_tickets (
			"_id"		INTEGER PRIMARY KEY,
			"DATE"		TEXT,
			"MARK"		TEXT,
			"MODEL"		TEXT,
			"TROUBLE"	TEXT,
			"TROUBLE1"	TEXT,
			"STREET"	TEXT,
			"HOUSE"		TEXT,
			"FLAT"		TEXT,
			"PHONE"		TEXT,
			"OTHER"		TEXT,
			"REPAIR"	TEXT,
			"REPAIR1"	TEXT,
			"WAR"		TEXT,
			"SUM"		TEXT,
			"NOTE"		TEXT
		);`
		execAny(path, sqlStatement)

		sqlStatement = `CREATE TABLE IF NOT EXISTS android_metadata 
			(locale TEXT DEFAULT 'en_US');`
		execAny(path, sqlStatement)

		log.Println("INFO: DB created!")
	}
}
