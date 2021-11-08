package stats

import (
	"reflect"
	"testing"

	"github.com/ManizhaM/bank/v2/pkg/types"
)

func TestCategoriesTotal(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 100000},
		{ID: 2, Category: "food", Amount: 200000},
		{ID: 3, Category: "auto", Amount: 300000},
		{ID: 4, Category: "auto", Amount: 400000},
		{ID: 5, Category: "fun", Amount: 500000},
	}
	expected := map[types.Category]types.Money{
		"auto": 800000,
		"food": 200000,
		"fun": 500000,
	}

	result := CategoriesTotal(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}	
}