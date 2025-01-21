package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight float64
	var userWeight float64
	fmt.Print("Enter height (cm): ")
	fmt.Scan(&userHeight)
	fmt.Print("Enter weight: ")
	fmt.Scan(&userWeight)
	var bodyMassIndex = userWeight / math.Pow(userHeight/100, 2)
	//fmt.Print("Your Body Mass Index: ")
	//fmt.Print(bodyMassIndex)
	//fmt.Print(int64(bodyMassIndex))
	//fmt.Printf("Your Body Mass Index: %v", bodyMassIndex)
	//fmt.Print("Your Body Mass Index: ", bodyMassIndex)
	fmt.Printf("Your Body Mass Index: %.2f", bodyMassIndex)
}
