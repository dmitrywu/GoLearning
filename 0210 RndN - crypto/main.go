package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
)

func main() {
	// Генерируем случайное число от 0 до 100

	n, err := rand.Int(rand.Reader, big.NewInt(100))
	if err != nil {
		log.Fatalf("Ошибка генерации случайного числа: %s", err.Error())
	}

	fmt.Println("Случайное число: ", n.Int64())

}
