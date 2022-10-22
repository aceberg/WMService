package db

import (
	"fmt"
	"log"
)

func ListInsert(path string, table string, name string) {

	sqlStatement := `INSERT INTO %s (NAME) VALUES ('%s');`

	sqlStatement = fmt.Sprintf(sqlStatement, table, name)

	execAny(path, sqlStatement)

	log.Println("INFO: added to list", table, "name", name)
}

func ListDelete(path string, table string, name string) {

	sqlStatement := `DELETE FROM %s WHERE NAME = '%s';`

	sqlStatement = fmt.Sprintf(sqlStatement, table, name)

	execAny(path, sqlStatement)

	log.Println("INFO: added to list", table, "name", name)
}
