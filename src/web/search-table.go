package web

import (
	// "fmt"
	. "github.com/aceberg/WMService/models"
	"log"
	"reflect"
	"strings"
)

func inString(str1 string, str2 string) bool {
	return strings.Contains(
		strings.ToLower(str1),
		strings.ToLower(str2),
	)
}

func searchTable(itemList []Item, field, search string) []Item {
	var foundList []Item

	foundList = []Item{}
	for _, item := range itemList {
		switch field {
		case "DATE":
			if inString(item.Date, search) {
				foundList = append(foundList, item)
			}
		case "MARK":
			if inString(item.Mark, search) {
				foundList = append(foundList, item)
			}
		case "MODEL":
			if inString(item.Model, search) {
				foundList = append(foundList, item)
			}
		case "TROUBLE":
			if inString(item.Trouble, search) {
				foundList = append(foundList, item)
			}
		case "TROUBLE1":
			if inString(item.Trouble1, search) {
				foundList = append(foundList, item)
			}
		case "STREET":
			if inString(item.Street, search) {
				foundList = append(foundList, item)
			}
		case "HOUSE":
			if inString(item.House, search) {
				foundList = append(foundList, item)
			}
		case "FLAT":
			if inString(item.Flat, search) {
				foundList = append(foundList, item)
			}
		case "PHONE":
			if inString(item.Phone, search) {
				foundList = append(foundList, item)
			}
		case "OTHER":
			if inString(item.Other, search) {
				foundList = append(foundList, item)
			}
		case "REPAIR":
			if inString(item.Repair, search) {
				foundList = append(foundList, item)
			}
		case "REPAIR1":
			if inString(item.Repair1, search) {
				foundList = append(foundList, item)
			}
		case "WAR":
			if inString(item.War, search) {
				foundList = append(foundList, item)
			}
		case "SUM":
			if inString(item.Sum, search) {
				foundList = append(foundList, item)
			}
		case "NOTE":
			if inString(item.Note, search) {
				foundList = append(foundList, item)
			}
		case "ANY":
			flag := false
			structValues := reflect.ValueOf(item)
			for i := 1; i < structValues.NumField(); i++ {
				val := structValues.Field(i).Interface().(string)

				if inString(val, search) {
					flag = true
				}
			}
			if flag {
				foundList = append(foundList, item)
			}
		default:
			log.Println("ERROR: wrong seach field")
		}
	}
	return foundList
}
