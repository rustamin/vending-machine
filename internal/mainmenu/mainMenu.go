package mainmenu

import (
	"fmt"

	"github.com/rustamin/vending-machine/internal/amount"
	"github.com/rustamin/vending-machine/internal/change"
	"github.com/rustamin/vending-machine/internal/good"
	"github.com/rustamin/vending-machine/internal/helper"
	"github.com/rustamin/vending-machine/internal/item"
	"github.com/rustamin/vending-machine/internal/returngate"
	"github.com/rustamin/vending-machine/model"
)

func Menu(goods []*model.Good, coins []*model.Coin, balance *int, chooseItem []model.Item, returnCoins []model.Coin) {

	helper.Line()
	fmt.Println(amount.List(balance))
	change.List(coins)
	returngate.List(returnCoins)
	good.List(goods)
	item.List(chooseItem)
	helper.Line()

	input := make([]int, 2)
	fmt.Print("Enter text: ")
	fmt.Scanln(&input[0], &input[1])

	if (input)[0] == 1 {
		coins, err := amount.UpdateAmountAndCoin(coins, balance, (input)[1])
		if err != nil {
			fmt.Println(err)
		}
		Menu(goods, coins, balance, chooseItem, returnCoins)
	} else if (input)[0] == 2 {
		chooseItem, err := item.UpdateChooseItem(chooseItem, goods, (input)[1], balance, coins)
		if err != nil {
			fmt.Println(err)
		}
		Menu(goods, coins, balance, chooseItem, returnCoins)
	} else if (input)[0] == 3 {
		// GET ITEM
		chooseItem, goods, err := item.GetChooseItem(chooseItem, goods, balance, coins)
		if err != nil {
			fmt.Println(err)
		}
		Menu(goods, coins, balance, chooseItem, returnCoins)
	} else if (input)[0] == 4 {
		coins, returnCoins := returngate.UpdateReturnGate(balance, coins)
		Menu(goods, coins, balance, chooseItem, returnCoins)
	} else if (input)[0] == 5 {
		returnCoins = returngate.GetReturn(returnCoins)
		Menu(goods, coins, balance, chooseItem, returnCoins)
	}
}
