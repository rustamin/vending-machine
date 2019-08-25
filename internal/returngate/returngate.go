package returngate

import (
	"github.com/rustamin/vending-machine/model"
)

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
