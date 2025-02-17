package main

import "fmt"

// iota - используется для констант, значение увеличивается на 1, в другой константной группе значение начинается снова с 0

/*
const (
	A = iota + 5 // начинаем перечисление с 5
	B
	C
	D
)*/

const (
	A = iota * 2 // Можно и с умножением значения каждого элемента
	B
	C
	D
)

const (
	_ = iota // можно пропускать элементы. Если мы пропускаем - ставим нижнее подчёркивание.
	K
	_
	L
)

const (
	Red = iota
	Green
	Blue
)

const (
	Orange = iota
	Yellow
	Violet
)

func main() {
	fmt.Println(Red, Green, Blue)
	fmt.Println(Orange, Yellow, Violet)
	fmt.Println(A, B, C, D)
	fmt.Println(K, L)

}
