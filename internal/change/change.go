package change

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/rustamin/vending-machine/model"
)

func List(coins []*model.Coin) {
	list := "[Change]"
	for i, elem := range coins {
		validCoin := map[int]bool{
			10:  true,
			100: true,
		}
		if !validCoin[elem.Nominal] {
			continue
		}
		var wordingChange string
		if elem.Total == 0 {
			wordingChange = "No change"
		} else {
			wordingChange = "Change"
		}

		if i == 0 {
			list += "           " + strconv.Itoa(elem.Nominal) + "      " + wordingChange
		} else {
			list += "\n                   " + strconv.Itoa(elem.Nominal) + "     " + wordingChange
		}
	}

	fmt.Println(list)
}

func ValidateCanDoChange(chooseItem []model.Item, balance *int, coins []*model.Coin) error {

	totalChange := totalChange(chooseItem, balance)

	if totalChange == 0 {
		return nil
	}

	if coins[0].Total < 9 || coins[2].Total < 4 {
		return errors.New("Sorry cannot choose/get item because coin return change bellow threshold. Please put in the exact amount of coin")
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
