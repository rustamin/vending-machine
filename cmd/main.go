package main

import (
	"github.com/rustamin/vending-machine/internal/mainmenu"
	"github.com/rustamin/vending-machine/model"
)

func main() {
	coins := make([]*model.Coin, 0)
	coins = append(coins, &model.Coin{
		Nominal: 10,
		Total:   1,
	})
	coins = append(coins, &model.Coin{
		Nominal: 50,
		Total:   0,
	})
	coins = append(coins, &model.Coin{
		Nominal: 100,
		Total:   0,
	})
	coins = append(coins, &model.Coin{
		Nominal: 500,
		Total:   0,
	})

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
	chooseItem := make([]model.Item, 0)

	mainmenu.Menu(goods, coins, &balance, chooseItem)
}
