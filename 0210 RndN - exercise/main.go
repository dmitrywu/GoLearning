package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	//randomNum := rand.IntN(100) // [0, 100)
	//fmt.Println(randomNum)

	/*
		Вам необходимо написать код внутри функции main. Для выполнения задания вам уже подготовлены все необходимые импорты и переменные.

		Ваша задача - создать переменную random типа int, которая будет содержать случайное целое число в диапазоне от min до max включительно
		 (данные переменные уже заданы и существуют в коде, обе переменные тип int).
	*/

	min := 1
	max := 10
	max++
	fmt.Println(max)
	var random int
	random = rand.IntN(max-min) + min
	fmt.Println(random)
}
