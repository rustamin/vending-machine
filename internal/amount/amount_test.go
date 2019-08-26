package amount

import (
	"fmt"
	"testing"

	"github.com/rustamin/vending-machine/model"
)

func TestList(t *testing.T) {
	tables := []struct {
		x int
		n string
	}{
		{100, "[Input amount]     100 JPY"},
	}

	for _, table := range tables {
		wording := List(&table.x)
		if wording != table.n {
			t.Errorf("Wording input amount of %d was incorrect, got: %s, want: %s.", table.x, wording, table.n)
		}
	}
}

func TestDeductBalance(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{100, 10, 90},
		{150, 10, 140},
	}

	for _, table := range tables {
		total := DeductBalance(&table.x, table.y)
		if *total != table.n {
			t.Errorf("Deduct of (%d-%d) was incorrect, got: %d, want: %d.", table.x, table.y, *total, table.n)
		}
	}
}

func TestUpdateBalance(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{100, 10, 110},
		{150, 10, 160},
	}

	for _, table := range tables {
		total := updateBalance(&table.x, table.y)
		if *total != table.n {
			t.Errorf("Update of (%d-%d) was incorrect, got: %d, want: %d.", table.x, table.y, *total, table.n)
		}
	}
}

func TestUpdateCoin(t *testing.T) {

	coins := make([]*model.Coin, 0)
	coins = append(coins, &model.Coin{
		Nominal: 10,
		Total:   10,
	})
	coins = append(coins, &model.Coin{
		Nominal: 50,
		Total:   0,
	})
	coins = append(coins, &model.Coin{
		Nominal: 100,
		Total:   10,
	})
	coins = append(coins, &model.Coin{
		Nominal: 500,
		Total:   0,
	})

	resultCoins := make([]model.Coin, 0)
	resultCoins = append(resultCoins, model.Coin{
		Nominal: 10,
		Total:   10,
	})
	resultCoins = append(resultCoins, model.Coin{
		Nominal: 50,
		Total:   0,
	})
	resultCoins = append(resultCoins, model.Coin{
		Nominal: 100,
		Total:   10,
	})
	resultCoins = append(resultCoins, model.Coin{
		Nominal: 500,
		Total:   0,
	})

	resultCoins[0].Total = resultCoins[0].Total + 1

	resultCoins1 := make([]model.Coin, len(resultCoins))
	copy(resultCoins1, resultCoins)
	resultCoins1[1].Total = resultCoins1[1].Total + 1

	resultCoins2 := make([]model.Coin, len(resultCoins))
	copy(resultCoins2, resultCoins)
	resultCoins2[2].Total = resultCoins2[2].Total + 1

	resultCoins3 := make([]model.Coin, len(resultCoins))
	copy(resultCoins3, resultCoins)
	resultCoins3[3].Total = resultCoins3[3].Total + 1

	tables := []struct {
		x []*model.Coin
		y int
		n []model.Coin
	}{
		{coins, 10, resultCoins},
		{coins, 50, resultCoins1},
		{coins, 100, resultCoins2},
		{coins, 500, resultCoins3},
	}

	for i, table := range tables {
		result := updateCoin(table.x, table.y)

		fmt.Println("testtt")
		if (result)[i].Total != (table.n)[i].Total {
			t.Errorf("Iterate: %d Update total coins of %d was incorrect, got total: %d, want: %d.", i, (table.x)[i].Nominal, (result)[i].Total, (table.n)[i].Total)
		}
	}
}

func TestValidateCoin(t *testing.T) {
	coins := make([]*model.Coin, 0)
	coins = append(coins, &model.Coin{
		Nominal: 10,
		Total:   0,
	})
	coins = append(coins, &model.Coin{
		Nominal: 50,
		Total:   0,
	})
	coins = append(coins, &model.Coin{
		Nominal: 100,
		Total:   10,
	})
	coins = append(coins, &model.Coin{
		Nominal: 500,
		Total:   0,
	})

	tables := []struct {
		x []*model.Coin
		y int
		n []*model.Coin
		m error
	}{
		{coins, 10, coins, nil},
		{coins, 20, coins, nil},
		{coins, 100, coins, nil},
	}

	for i, table := range tables {
		result, err := validateCoin(table.x, table.y)
		if result[i].Total != table.n[i].Total {
			t.Errorf("Iterate: %d Update total coins of %d was incorrect, got total: %d, want: %d.", i, (table.x)[i].Nominal, (result)[i].Total, (table.n)[i].Total)
		}

		Expected := "Please insert 10, 50, 100, 500 JPY"

		if i == 1 && err.Error() != Expected {
			t.Errorf("Iterate: %d Error validate coin of insrted coin %d was incorrect, got actual: %v, want: %v.", i, table.y, err, Expected)
		} else if i == 2 && err.Error() != "Running out 10 JPY. Please insert 10 JPY" {
			t.Errorf("Iterate: %d Error validate coin of insrted coin %d was incorrect, got actual: %v, want: %v.", i, table.y, err, Expected)
		}

	}
}

func TestUpdateAmountAndCoin(t *testing.T) {

	coins := make([]*model.Coin, 0)
	coins = append(coins, &model.Coin{
		Nominal: 10,
		Total:   10,
	})
	coins = append(coins, &model.Coin{
		Nominal: 50,
		Total:   0,
	})
	coins = append(coins, &model.Coin{
		Nominal: 100,
		Total:   10,
	})
	coins = append(coins, &model.Coin{
		Nominal: 500,
		Total:   0,
	})

	resultCoins := make([]model.Coin, 0)
	resultCoins = append(resultCoins, model.Coin{
		Nominal: 10,
		Total:   10,
	})
	resultCoins = append(resultCoins, model.Coin{
		Nominal: 50,
		Total:   0,
	})
	resultCoins = append(resultCoins, model.Coin{
		Nominal: 100,
		Total:   10,
	})
	resultCoins = append(resultCoins, model.Coin{
		Nominal: 500,
		Total:   0,
	})

	resultCoins[0].Total = resultCoins[0].Total + 1

	balance := 0

	tables := []struct {
		x []*model.Coin
		y *int
		z int
		n []model.Coin
	}{
		{coins, &balance, 10, resultCoins},
	}

	for i, table := range tables {
		result, err := UpdateAmountAndCoin(table.x, table.y, table.z)
		Expected := "Running out 10 JPY. Please insert 10 JPY"

		fmt.Println("testtt")
		if (result)[i].Total != (table.n)[i].Total {
			t.Errorf("Iterate: %d Update total coins of %d was incorrect, got total: %d, want: %d.", i, (table.x)[i].Nominal, (result)[i].Total, (table.n)[i].Total)
		} else if i == 1 && err.Error() != Expected {
			t.Errorf("Iterate: %d Error validate coin of insrted coin %d was incorrect, got actual: %v, want: %v.", i, table.y, err, Expected)
		}
	}
}
