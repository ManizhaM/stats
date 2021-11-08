package stats

import (
	"github.com/ManizhaM/bank/v2/pkg/types"
	"reflect"
	"testing"
)

func TestFilterByCategory_nil(t *testing.T) {
	var payments []types.Payment
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}	
}

func TestFilterByCategory_empty(t *testing.T) {
	 payments := []types.Payment{}
	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}	
}

func TestFilterByCategory_notFound(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	} 

	result := FilterByCategory(payments, "mobile")

	if len(result) != 0 {
		t.Error("result len != 0")
	}
}

func TestFilterByCategory_foundOne(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	} 
	expected := []types.Payment{
		{ID: 2, Category: "food"},
	}

	result := FilterByCategory(payments, "food")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}

func TestFilterByCategory_foundMultiple(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 2, Category: "food"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
		{ID: 5, Category: "fun"},
	} 
	expected := []types.Payment{
		{ID: 1, Category: "auto"},
		{ID: 3, Category: "auto"},
		{ID: 4, Category: "auto"},
	}

	result := FilterByCategory(payments, "auto")

	if !reflect.DeepEqual(expected, result) {
		t.Error("invalid result")
	}
}


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

func TestCategoriesAvg(t *testing.T) {
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount: 100000},
		{ID: 2, Category: "food", Amount: 200000},
		{ID: 2, Category: "food", Amount: 100000},
		{ID: 3, Category: "auto", Amount: 300000},
		{ID: 4, Category: "auto", Amount: 200000},
		{ID: 5, Category: "fun", Amount: 500000},
	}
	expected := map[types.Category]types.Money{
		"auto": 200000,
		"food": 150000,
		"fun": 500000,
	}

	result := CategoriesAvg(payments)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}	
}


func TestPeriodsDynamic(t *testing.T) {
	first := map[types.Category]types.Money{
		"auto": 200000,
		"food": 150000,
		"fun": 500000,
	}

	second := map[types.Category]types.Money{
		"auto": 100000,
		"food": 200000,
		"fun": 500000,
	}
	expected := map[types.Category]types.Money{
		"auto": -100000,
		"food": 50000,
		"fun": 0,
	}

	result := PeriodsDynamic(first, second)

	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v", expected, result)
	}	
}