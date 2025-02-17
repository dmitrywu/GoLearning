package main

import (
	"fmt"
	"math"
)

func main() {

	// преобразование к float64
	var i int = 42
	var f float64 = float64(i)
	println(f) // разница в выводе между fmt.Println и println
	fmt.Println("Converted int to float: ", f)

	// преобразование к int
	var f2 float64 = 42.9
	var i2 int = int(f2) // дробная часть просто отбрасывается
	fmt.Println("Converted float to int: ", i2)

	// преобразование к int c округлением
	var f3 float64 = 42.9
	var i3 int = int(math.Round(f3)) // Округление
	fmt.Println("Converted float to int: ", i3)

}
