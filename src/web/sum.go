package web

import (
	. "github.com/aceberg/WMService/models"
	"strconv"
)

func getSum(itemList []Item) int {
	var sum, itemsum int

	sum = 0
	for _, item := range itemList {
		itemsum, _ = strconv.Atoi(item.Sum)
		sum = sum + itemsum
	}
	return sum
}
