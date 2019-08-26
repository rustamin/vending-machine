package item

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rustamin/vending-machine/model"
)

func TestList(t *testing.T) {
	chooseItem := make([]model.Item, 0)
	chooseItem = append(chooseItem, model.Item{
		Name:  "Test",
		Price: 100,
	})

	tables := []struct {
		chooseItem []model.Item
		result     string
	}{
		{chooseItem, "[Outlet]           Test"},
	}

	for _, table := range tables {
		wording := List(table.chooseItem)

		if !strings.Contains(wording, table.result) {
			t.Errorf("Wording was incorrect, got: %s, want: %s.", wording, table.result)
		}
	}
}

func TestUpdateChooseItem(t *testing.T) {

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

	chooseItem := make([]model.Item, 0)

	goods := make([]*model.Good, 0)
	goods = append(goods, &model.Good{
		Name:  "Canned coffee",
		Price: 100,
		Stock: 2,
	})
	goods = append(goods, &model.Good{
		Name:  "Canned coffee",
		Price: 100,
		Stock: 0,
	})

	balance := 0
	balance1 := 100

	tables := []struct {
		chooseItem []model.Item
		goods      []*model.Good
		input      int
		balance    *int
		coins      []*model.Coin
	}{
		{chooseItem, goods, 1, &balance, coins},
		{chooseItem, goods, 2, &balance1, coins},
		{chooseItem, goods, 1, &balance1, coins},
	}

	Expected := "Balance is not enough. Insert more coins"
	Expected1 := "Stock is empty. Please choose another good"

	for i, table := range tables {
		result, err := UpdateChooseItem(table.chooseItem, table.goods, table.input, table.balance, table.coins)

		fmt.Println("testtt")
		if i == 0 && err.Error() != Expected {
			t.Errorf("Iterate: %d Should return error balance not enough, got result: %s, want: %s.", i, err.Error(), Expected)
		} else if i == 1 && err.Error() != Expected1 {
			t.Errorf("Iterate: %d Should return error balance not enough, got result: %s, want: %s.", i, err.Error(), Expected1)
		} else if i == 2 && goods[0].Name != result[0].Name {
			fmt.Println(result)
			t.Errorf("Iterate: %d Should add to choose item, got result: %s, want: %s.", i, goods[0].Name, result[0].Name)
		}
	}
}

func TestGetChooseItem(t *testing.T) {
	chooseItem := make([]model.Item, 0)
	chooseItem = append(chooseItem, model.Item{
		Name:  "Test",
		Price: 100,
	})

	tables := []struct {
		chooseItem          []model.Item
		lenResultChooseItem int
	}{
		{chooseItem, 0},
	}

	for _, table := range tables {
		result := GetChooseItem(table.chooseItem)
		if len(result) != 0 {
			t.Errorf("Failed to empty choose item, got: %d, want: %d.", len(result), table.lenResultChooseItem)

		}
	}
}
