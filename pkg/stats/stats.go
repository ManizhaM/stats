package stats

import 	"github.com/ManizhaM/bank/v2/pkg/types"

//Avg Рассчитывает среднюю сумму платежа
func Avg(payments []types.Payment) types.Money{
	var sum types.Money = 0 
	count := 0 
	for _, payment := range payments {
		if(payment.Status != types.StatusFail){
		sum += payment.Amount 
		count++;
		}
	}
	sum = sum/types.Money(count)
	return sum
}

//TotalInCategory находит сумму покупок в определенной категории
func TotalInCategory(payments []types.Payment, category types.Category) types.Money{
	var sum types.Money = 0
	for _, payment := range payments {
		if(payment.Category == category && payment.Status != types.StatusFail){
			sum+=payment.Amount
		}
	}
	return sum
}


//FilterByCategory возвращает платежи в указанной категории
func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment{
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}

//CategoriesTotal возвращает сумму платежей по каждой категории
func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	
	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}
	return categories
}
//CategoriesAvg возвращает среднюю сумму платежа по каждой категории
func CategoriesAvg(payments []types.Payment)map[types.Category]types.Money{
	categories := map[types.Category]types.Money{}
	count := map[types.Category]types.Money{}
	for _, payment := range payments {
		categories[payment.Category]+=payment.Amount
		count[payment.Category]++;
	}
	for key, value := range categories {
		categories[key] = value/count[key]
	}
return categories
}