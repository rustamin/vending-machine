package good

import (
	"strings"
	"testing"

	"github.com/rustamin/vending-machine/model"
)

func TestList(t *testing.T) {
	goods := make([]*model.Good, 0)
	goods = append(goods, &model.Good{
		Name:  "Canned coffee",
		Price: 100,
		Stock: 2,
	})
	goods = append(goods, &model.Good{
		Name:  "Water PET bottle",
		Price: 100,
		Stock: 0,
	})

	tables := []struct {
		goods  []*model.Good
		result string
	}{
		{goods, "[Items for sale]   1. Canned coffee         100JPY"},
		{goods, "2. Water PET bottle      10JPY"},
	}

	for _, table := range tables {
		wording := List(table.goods)

		if !strings.Contains(wording, table.result) {
			t.Errorf("Wording change was incorrect, got: %s, want: %s.", wording, table.result)
		}
	}
}
