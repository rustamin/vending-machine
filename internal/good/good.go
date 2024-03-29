package good

import (
	"strconv"

	"github.com/rustamin/vending-machine/internal/helper"
	"github.com/rustamin/vending-machine/model"
)

func List(goods []*model.Good) string {
	list := "[Items for sale]"

	for i, elem := range goods {
		var number int
		number = i + 1
		item := strconv.Itoa(number) + ". " + elem.Name + helper.Padding(elem.Name) + strconv.Itoa(elem.Price) + "JPY"

		if i == 0 {
			list += "   " + item
		} else {
			list += "\n                   " + item
		}
		if elem.Stock == 0 {
			list += "   Sold out"
		}
	}
	return list
}
