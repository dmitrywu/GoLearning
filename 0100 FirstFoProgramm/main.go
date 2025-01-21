package main

import (
	"fmt"
	"math"
)

func main() {
	//	fmt.Printf("Hello world!")
	//var userHeight = 1.75
	//var userWeight = 117.5
	//var userHeight float64 = 1.75

	const bmiPow = 2 // Константы

	userHeight := 1.75     // При такой записи (:=) нельзя задать тип данных, он присваивается автоматически. И нельзя просто объявить переменную без присваивания значения
	var userWeight float64 // Если объявить переменную, но не присвоить значение, тип нужно задавать обязательно.
	userWeight = 117.5

	// Можно в одной строке задавать сразу несколько переменных и задать их тип, можно использовать := (при этом нельзя задать тип)
	//var user1, user2 int64 = 4,6

	var bodyMassIndex = userWeight / math.Pow(userHeight, 2)
	//var bodyMassIndex = float64(userWeight) / userHeight // С преобразование типа переменной к нужному float64()

	fmt.Print(bodyMassIndex)

}
