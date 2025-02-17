package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	//randomNum := rand.IntN(100) // [0, 100)
	//fmt.Println(randomNum)

	min := 1
	max := 10
	max++
	fmt.Println(max)
	var random int
	random = rand.IntN(max-min) + min
	fmt.Println(random)
}
