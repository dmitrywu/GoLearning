package main

import "fmt"

func main() {
	// byte is alias for uint8, ASCII
	var b byte = 200
	println(b, string(b))

	s := "hi"
	println(s[0], string(s[0]))
	println(s[1], string(s[1]))

	// rune is alias for int32, UniCode, одинарные кавычки, "это символ, а не просто число"
	//r := "😀"
	//println(r[0], string(r[0]), r)
	r := '😀'
	fmt.Println(r, string(r))
}
