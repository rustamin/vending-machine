package item

import (
	"errors"
	"fmt"

	"github.com/rustamin/vending-machine/model"
)

func List(chooseItem []model.Item) {
	list := "[Outlet]"
	for i, elem := range chooseItem {
		if i == 0 {
			list += "           " + elem.Name
		} else {
			list += "\n                   " + elem.Name
		}
	}

	fmt.Println(list)
}

func UpdateChooseItem(chooseItem []model.Item, goods []*model.Good, input int, balance *int) ([]model.Item, error) {

	item := new(model.Item)

	for i, elem := range goods {

		if i == input-1 {
			// TODO: validate if good is not found
			if *balance-elem.Price < 0 {
				return chooseItem, errors.New("Balance is not enough. Insert more coins")
			}

			// *balance = *balance - elem.Price
			chooseItem = append(chooseItem, model.Item{
				Name:  elem.Name,
				Price: elem.Price,
			})

			item.Name = elem.Name
			item.Price = elem.Price
		}
	}
	return chooseItem, nil
}
