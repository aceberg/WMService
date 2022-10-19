package db

import (
	"fmt"
	. "github.com/aceberg/WMService/models"
	"log"
)

func AddItem(path string, item Item) {

	sqlStatement := `INSERT INTO rs_tickets (DATE, MARK, MODEL, TROUBLE, TROUBLE1, STREET, HOUSE, FLAT, PHONE, OTHER, REPAIR, REPAIR1, WAR, SUM, NOTE)
	VALUES ('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s');`

	sqlStatement = fmt.Sprintf(sqlStatement, item.Date, item.Mark, item.Model, item.Trouble, item.Trouble1, item.Street, item.House, item.Flat, item.Phone, item.Other, item.Repair, item.Repair1, item.War, item.Sum, item.Note)

	execAny(path, sqlStatement)

	log.Println("INFO: added to table", item)
}
