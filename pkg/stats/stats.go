package stats

import 	"github.com/ManizhaM/bank/pkg/types"

//Avg Рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money{
	var sum types.Money = 0 
	count := 0 
	for _, payment := range payments {
		sum += payment.Amount 
		count++;
	}
	sum = sum/types.Money(count)
	return sum
}

//TotalInCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money{
	var sum types.Money = 0
	for _, payment := range payments {
		if(payment.Category == category){
			sum+=payment.Amount
		}
	}
	return sum
}