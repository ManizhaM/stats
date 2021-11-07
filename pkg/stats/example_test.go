package stats

import (
	"fmt"
	"github.com/ManizhaM/bank/v2/pkg/types"
)


// Тестирование функции Avg - Рассчитывает среднюю сумму платежа
func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount:1000,
			Status: types.StatusFail,
		},
		{
			Amount:2000,
			Status: types.StatusOk,
		},
		{
			Amount:3000,
			Status: types.StatusFail,
		},
	}
	fmt.Println(Avg(payments))
	//Output:
	//2000
}

// Тест на TotalInCategory - находит сумму покупок в определенной категории
func ExampleTotalInCategory(){
	payments := []types.Payment{
		{
			Amount:1000,
			Category: "car",
			Status: types.StatusOk,
		},
		{
			Amount:2000,
			Category: "dress",
			Status: types.StatusOk,
		},
		{
			Amount:3000,
			Category: "dress",
			Status: types.StatusFail,
		},
	}
	fmt.Println(TotalInCategory(payments, "car"))
	//Output:
	//1000
}
