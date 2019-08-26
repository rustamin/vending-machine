package helper

import (
	"testing"
)

func TestPadding(t *testing.T) {
	tables := []struct {
		str    string
		result int
	}{
		{"test", 18}, //
		{"double test", 11},
	}

	for _, table := range tables {
		result := Padding(table.str)
		if len(result) != table.result {
			t.Errorf("Padding length was incorrect, got: %d, want: %d.", len(result), table.result)
		}
	}
}

func TestLine(t *testing.T) {
	tables := []struct {
		result string
	}{
		{"-------------------------------------------------------"},
	}

	for _, table := range tables {
		result := Line()
		if result != table.result {
			t.Errorf("Line was incorrect, got: %s, want: %s.", result, table.result)
		}
	}
}
