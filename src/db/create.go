package db

import (
// "log"
// "os"
)

func CreateDB(path string) {

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

	sqlStatement = `CREATE TABLE IF NOT EXISTS rs_mark (
		"_id"		INTEGER PRIMARY KEY,
		"NAME"		TEXT
	);`
	execAny(path, sqlStatement)

	sqlStatement = `CREATE TABLE IF NOT EXISTS rs_trouble (
		"_id"		INTEGER PRIMARY KEY,
		"NAME"		TEXT
	);`
	execAny(path, sqlStatement)

	sqlStatement = `CREATE TABLE IF NOT EXISTS rs_repair (
		"_id"		INTEGER PRIMARY KEY,
		"NAME"		TEXT
	);`
	execAny(path, sqlStatement)

	// log.Println("INFO: DB created!")
}
