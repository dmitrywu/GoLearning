package main

import (
	"fmt"
	"math/rand/v2"
	// crypto/rand
)

func main() {
	//randomNum := rand.IntN(100) // [0, 100)
	//fmt.Println(randomNum)

	// [ - включая, ) - не включая

	min := 10
	max := 50

	randomNum := rand.IntN(max-min) + 10 // [10, 50)
	fmt.Println(randomNum)
}
