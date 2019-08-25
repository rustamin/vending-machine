package good

import (
	"fmt"
	"strconv"

	"github.com/rustamin/vending-machine/internal/helper"
	"github.com/rustamin/vending-machine/model"
	// "github.com/rustamin/vending-machine/model"
)

func List(goods []*model.Good) {
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

	}

	fmt.Println(list)

}
