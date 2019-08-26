package amount

import (
	"errors"
	"strconv"

	"github.com/rustamin/vending-machine/model"
)

func List(amount *int) string {
	msgAmount := "[Input amount]     " + strconv.Itoa(*amount) + " JPY"
	return msgAmount
}

func UpdateAmountAndCoin(coins []*model.Coin, balance *int, input int) ([]*model.Coin, error) {
	coins, err := validateCoin(coins, input)
	if err != nil {
		return coins, err
	}

	coins = updateCoin(coins, input)
	updateBalance(balance, input)
	return coins, nil
}

// TODO: simplify update balance and deduct balance
func updateBalance(balance *int, input int) *int {
	*balance = *balance + input
	return balance
}

func DeductBalance(balance *int, price int) *int {
	*balance = *balance - price
	return balance
}

func updateCoin(coins []*model.Coin, input int) []*model.Coin {
	if input == 10 {
		(coins)[0].Total++
	} else if input == 50 {
		(coins)[1].Total++
	} else if input == 100 {
		(coins)[2].Total++
	} else if input == 500 {
		(coins)[3].Total++
	}

	return coins
}

func validateCoin(coins []*model.Coin, input int) ([]*model.Coin, error) {
	validCoin := map[int]bool{
		10:  true,
		50:  true,
		100: true,
		500: true,
	}
	if !validCoin[input] {
		return coins, errors.New("Please insert 10, 50, 100, 500 JPY")
	}

	if (coins)[0].Total == 0 && input != 10 {
		return coins, errors.New("Running out 10 JPY. Please insert 10 JPY")
	}
	return coins, nil
}
