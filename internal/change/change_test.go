package change

import (
	"errors"
	"strings"
	"testing"

	"github.com/rustamin/vending-machine/model"
)

func TestList(t *testing.T) {
	coins := make([]*model.Coin, 0)
	coins = append(coins, &model.Coin{
		Nominal: 10,
		Total:   1,
	})
	coins = append(coins, &model.Coin{
		Nominal: 100,
		Total:   0,
	})

	tables := []struct {
		x []*model.Coin
		n string
	}{
		{coins, "[Change]           10      Change"},
		{coins, "100     No change"},
	}

	for _, table := range tables {
		wording := List(table.x)

		if !strings.Contains(wording, table.n) {
			t.Errorf("Wording change was incorrect, got: %s, want: %s.", wording, table.n)
		}
	}
}

func TestValidateCanDoChange(t *testing.T) {
	chooseItem := make([]model.Item, 0)
	chooseItem = append(chooseItem, model.Item{
		Name:  "Test",
		Price: 100,
	})

	coins := make([]*model.Coin, 0)
	coins = append(coins, &model.Coin{
		Nominal: 10,
		Total:   1,
	})
	coins = append(coins, &model.Coin{
		Nominal: 50,
		Total:   1,
	})
	coins = append(coins, &model.Coin{
		Nominal: 100,
		Total:   0,
	})

	balance := 100
	balance1 := 150

	tables := []struct {
		chooseItem []model.Item
		balance    *int
		coins      []*model.Coin
		err        error
	}{
		{chooseItem, &balance, coins, nil},
		{chooseItem, &balance1, coins, errors.New("Sorry cannot choose/get item because coin return change bellow threshold. Please put in the exact amount of coin")},
	}

	for i, table := range tables {
		err := ValidateCanDoChange(table.chooseItem, table.balance, table.coins)
		if i == 0 && err != table.err {
			t.Errorf("Validate can do change was incorrect, got: %s, want: %s.", err, table.err)
		} else if i == 1 && err.Error() != "Sorry cannot choose/get item because coin return change bellow threshold. Please put in the exact amount of coin" {
			t.Errorf("Validate can do change was incorrect, got: %s, want: %s.", err, table.err)
		}
	}
}
