package db

import (
	"fmt"
	. "github.com/aceberg/WMService/models"
	"log"
)

func UpdateItem(path string, item Item) {

	item = quoteItem(item)

	sqlStatement := `UPDATE rs_tickets SET DATE = '%s', MARK = '%s', MODEL = '%s', TROUBLE = '%s', TROUBLE1 = '%s', STREET = '%s', HOUSE = '%s', FLAT = '%s', PHONE = '%s', OTHER = '%s', REPAIR = '%s', REPAIR1 = '%s', WAR = '%s', SUM = '%s', NOTE = '%s' WHERE _id = '%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, item.Date, item.Mark, item.Model, item.Trouble, item.Trouble1, item.Street, item.House, item.Flat, item.Phone, item.Other, item.Repair, item.Repair1, item.War, item.Sum, item.Note, item.Id)

	execAny(path, sqlStatement)

	log.Println("INFO: updated", item)
}

func DeleteItem(path string, id int) {

	sqlStatement := `DELETE FROM rs_tickets WHERE _id = '%d';`

	sqlStatement = fmt.Sprintf(sqlStatement, id)

	execAny(path, sqlStatement)

	log.Println("INFO: deleted id =", id)
}
