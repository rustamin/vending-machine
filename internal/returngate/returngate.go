package returngate

import (
	"fmt"
	"strconv"

	"github.com/rustamin/vending-machine/model"
)

func List(coinsToReturn []model.Coin) {
	list := "[Return gate]"

	if len(coinsToReturn) == 0 {
		list += "      Empty"
	}

	for i, elem := range coinsToReturn {
		item := strconv.Itoa(elem.Nominal) + " JPY"

		if i == 0 {
			list += "      " + item
		} else {
			list += "\n                   " + item
		}
	}
	fmt.Println(list)
}

func UpdateReturnGate(balance *int, coins []*model.Coin) ([]*model.Coin, []model.Coin) {

	coinsToReturn := make([]model.Coin, 0)

	reserveCoin := make([]*model.Coin, len(coins))
	copy(reserveCoin, coins)

	for i, j := 0, len(reserveCoin)-1; i < j; i, j = i+1, j-1 {
		reserveCoin[i], reserveCoin[j] = reserveCoin[j], reserveCoin[i]
	}

	for _, elem := range reserveCoin {
		for *balance >= elem.Nominal && elem.Total > 0 {
			*balance = *balance - elem.Nominal
			elem.Total = elem.Total - 1
			coinsToReturn = append(coinsToReturn, model.Coin{
				Nominal: elem.Nominal,
				Total:   1,
			})
		}
	}

	// reserve back coins
	for i, j := 0, len(reserveCoin)-1; i < j; i, j = i+1, j-1 {
		reserveCoin[i], reserveCoin[j] = reserveCoin[j], reserveCoin[i]
	}
	coins = reserveCoin

	return coins, coinsToReturn
}

func GetReturn(returnCoins []model.Coin) []model.Coin {
	returnCoins = make([]model.Coin, 0)
	return returnCoins
}
