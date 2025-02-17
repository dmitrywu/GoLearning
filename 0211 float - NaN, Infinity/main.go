package main

import (
	"fmt"
	"math"
)

func main() {
	// NaN - not a number - не число,
	// используется, когда мы не можем какие-то числовые операции сделать,
	// либо когда результат не известен
	//
	// +Infinity - бесконечность. Когда мы используем слишком большое число
	// -Infinity - ... или слишком маленькое число
	// это относится к типу float

	// zeroFloat64 := 0 // Это Int
	zeroFloat64 := 0.0 // Это float
	result2 := 1.0 / zeroFloat64
	result4 := -1.0 / zeroFloat64
	result3 := zeroFloat64 / zeroFloat64
	nan := math.NaN() // Создаём переменную со значением NaN
	posInf := math.Inf(1)
	negInf := math.Inf(-1)

	fmt.Println("+1/0: ", result2)
	fmt.Println("-1/0: ", result4)
	fmt.Println("0/0: ", result3)
	fmt.Println("NaN: ", nan)
	fmt.Println("Inf: ", posInf)
	fmt.Println("This is NaN? : ", math.IsNaN(nan))                    // Проверяем, NaN или что это?
	fmt.Println("This is +Infinity? : ", math.IsInf(posInf, 1))        // Проверяем, +Infinity или что это?
	fmt.Println("This is -Infinity? : ", math.IsInf(negInf, -1))       // Проверяем, -Infinity или что это?
	fmt.Println("This is Infinity(posInf)? : ", math.IsInf(posInf, 0)) // Проверяем, Infinity или что это?
	fmt.Println("This is Infinity(negInf)? : ", math.IsInf(negInf, 0)) // Проверяем, Infinity или что это?

}
