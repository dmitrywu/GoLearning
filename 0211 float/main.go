package main

import (
	"fmt"
)

func main() {
	var a float32
	var b float32 = 3.1415
	var c float32 = 3.1415
	// можно использовать нижнее подчёркивание, чтобы разделить блоки цифр, просто для удобства восприятия
	f := 312_321.145_464_611

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(f)
}
