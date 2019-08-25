package item

import (
	"fmt"

	"github.com/rustamin/vending-machine/model"
)

func UpdateChooseItem(chooseItem []model.Item, goods []*model.Good, input int) {

	fmt.Println("LIST GOODS")
	fmt.Println(goods)
	item := new(model.Item)

	for i, elem := range goods {

		if i == input-1 {

			chooseItem = append(chooseItem, model.Item{
				Name:  elem.Name,
				Price: elem.Price,
			})

			item.Name = elem.Name
			item.Price = elem.Price
		}
	}
	// fmt.Println("ITEM")
	// fmt.Println(*item)

	// *chooseItem = *amount + input
	// goods = append(goods, &model.Good{
	// 	Name:  "Sport drinks",
	// 	Price: 150,
	// 	Stock: 0,
	// })
	// chooseItem = append(chooseItem, &item)
	// chooseItem = append(chooseItem, item)
	// *chooseItem = append(*chooseItem, &model.Item{
	// 	Name:  item.Name,
	// 	Price: item.Price,
	// })
	fmt.Println("CHOOSE ITEM")
	fmt.Println(chooseItem)
	// return *amount
}
