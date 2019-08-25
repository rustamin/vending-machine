package mainmenu

import (
	"fmt"

	"github.com/rustamin/vending-machine/internal/amount"
	"github.com/rustamin/vending-machine/internal/good"
	"github.com/rustamin/vending-machine/internal/helper"
	"github.com/rustamin/vending-machine/internal/item"
	"github.com/rustamin/vending-machine/model"
)

func Menu(goods []*model.Good, balance *int, chooseItem []model.Item) {

	helper.Line()
	amount.Menu(balance)
	good.List(goods)
	helper.Line()

	input := make([]int, 2)
	fmt.Print("Enter text: ")
	fmt.Scanln(&input[0], &input[1])

	if (input)[0] == 1 {
		amount.UpdateAmount(balance, (input)[1])
		fmt.Println("BALANCE IN MAIN MENU")
		fmt.Println(balance)
		Menu(goods, balance, chooseItem)
	}
	if (input)[0] == 2 {
		item.UpdateChooseItem(chooseItem, goods, (input)[1])
		fmt.Println("CHOSE ITEM IN MAIN MENU")
		fmt.Println(chooseItem)
		Menu(goods, balance, chooseItem)
	}

	fmt.Println("balance adlah")
	fmt.Println(balance)
}

// func switcher(input *[]int, goods []*model.Good) {
// 	var balance int

// 	if (*input)[0] == 1 {
// 		balance = (*input)[1]
// 		Menu(goods)
// 	}
// 	fmt.Println("balance adlah")
// 	fmt.Println(balance)

// }
