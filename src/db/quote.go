package db

import (
	. "github.com/aceberg/WMService/models"
	"strings"
)

func quoteStr(str string) string {
	return strings.ReplaceAll(str, "'", "''")
}

func quoteItem(item Item) Item {
	var newItem Item

	newItem.Id = item.Id
	newItem.Date = strings.ReplaceAll(item.Date, "'", "''")
	newItem.Mark = strings.ReplaceAll(item.Mark, "'", "''")
	newItem.Model = strings.ReplaceAll(item.Model, "'", "''")
	newItem.Trouble = strings.ReplaceAll(item.Trouble, "'", "''")
	newItem.Trouble1 = strings.ReplaceAll(item.Trouble1, "'", "''")
	newItem.Street = strings.ReplaceAll(item.Street, "'", "''")
	newItem.House = strings.ReplaceAll(item.House, "'", "''")
	newItem.Flat = strings.ReplaceAll(item.Flat, "'", "''")
	newItem.Phone = strings.ReplaceAll(item.Phone, "'", "''")
	newItem.Other = strings.ReplaceAll(item.Other, "'", "''")
	newItem.Repair = strings.ReplaceAll(item.Repair, "'", "''")
	newItem.Repair1 = strings.ReplaceAll(item.Repair1, "'", "''")
	newItem.War = strings.ReplaceAll(item.War, "'", "''")
	newItem.Sum = strings.ReplaceAll(item.Sum, "'", "''")
	newItem.Note = strings.ReplaceAll(item.Note, "'", "''")

	return newItem
}
