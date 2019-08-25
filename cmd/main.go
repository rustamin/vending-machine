package main

import (
	"github.com/rustamin/vending-machine/internal/mainmenu"
	"github.com/rustamin/vending-machine/model"
)

func main() {
	goods := make([]*model.Good, 0)
	goods = append(goods, &model.Good{
		Name:  "Canned coffee",
		Price: 120,
		Stock: 2,
	})
	goods = append(goods, &model.Good{
		Name:  "Water PET bottle",
		Price: 100,
		Stock: 1,
	})
	goods = append(goods, &model.Good{
		Name:  "Sport drinks",
		Price: 150,
		Stock: 0,
	})

	var balance int
	// var chooseItem []string
	chooseItem := make([]model.Item, 10)

	mainmenu.Menu(goods, &balance, chooseItem)
}
