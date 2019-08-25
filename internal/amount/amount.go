package amount

import (
	"fmt"
	"strconv"
)

func Menu(amount *int) {
	msgAmount := "[Input amount]     " + strconv.Itoa(*amount) + " JPY"
	fmt.Println(msgAmount)
}

func UpdateAmount(amount *int, input int) int {
	*amount = *amount + input
	return *amount

}
