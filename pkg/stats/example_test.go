package stats

import (
	"fmt"
	"github.com/ManizhaM/bank/pkg/types"
)


// Тестирование функции Avg - Рассчитывает среднюю сумму платежа
func ExampleAvg() {
	payments := []types.Payment{
		{
			Amount:1000,
		},
		{
			Amount:2000,
		},
		{
			Amount:3000,
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
		},
		{
			Amount:2000,
			Category: "dress",
		},
		{
			Amount:3000,
			Category: "dress",
		},
	}
	fmt.Println(TotalInCategory(payments, "car"))
	//Output:
	//1000
}
