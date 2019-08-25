package item

import (
	"errors"
	"fmt"

	"github.com/rustamin/vending-machine/internal/amount"
	"github.com/rustamin/vending-machine/internal/change"
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

func UpdateChooseItem(chooseItem []model.Item, goods []*model.Good, input int, balance *int, coins []*model.Coin) ([]model.Item, error) {
	item := new(model.Item)

	for i, elem := range goods {

		if i == input-1 {
			// TODO: validate if good is not found
			if *balance-elem.Price < 0 {
				return chooseItem, errors.New("Balance is not enough. Insert more coins")
			} else if elem.Stock == 0 {
				return chooseItem, errors.New("Stock is empty. Please choose another good")
			}

			tempChooseItem := make([]model.Item, 0)
			tempChooseItem = append(tempChooseItem, model.Item{
				Name:  elem.Name,
				Price: elem.Price,
			})

			err := change.ValidateCanDoChange(tempChooseItem, balance, coins)
			if err != nil {
				return chooseItem, err
			}

			// *balance = *balance - elem.Price
			amount.DeductBalance(balance, elem.Price)
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

func GetChooseItem(chooseItem []model.Item, goods []*model.Good, balance *int, coins []*model.Coin) ([]model.Item, []*model.Good, error) {
	// err := change.ValidateCanDoChange(chooseItem, balance, coins)
	// if err != nil {
	// 	return chooseItem, goods, err
	// }

	// empty outlet
	chooseItem = make([]model.Item, 0)
	return chooseItem, goods, nil
}
