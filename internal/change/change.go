package change

import (
	"errors"
	"fmt"

	"github.com/rustamin/vending-machine/model"
)

// func DoChange() {

// }

func ValidateCanDoChange(chooseItem []model.Item, balance *int, coins []*model.Coin) error {
	// validate return, if not posible throw error
	fmt.Println("chooseItemchooseItem")
	fmt.Println(chooseItem)

	totalChange := totalChange(chooseItem, balance)

	if totalChange == 0 {
		return nil
	}

	if coins[0].Total < 9 || coins[2].Total < 4 {
		return errors.New("Sorry cannot get item because coin return change bellow threshold. Please put in the exact amount of coin")
	}

	return nil
}

func totalChange(chooseItem []model.Item, balance *int) int {
	totalPriceChooseItem := 0

	for _, elem := range chooseItem {
		totalPriceChooseItem += elem.Price
	}

	totalChange := *balance - totalPriceChooseItem
	return totalChange
}
