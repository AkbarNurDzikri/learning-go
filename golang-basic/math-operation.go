package main

import "fmt"

func main() {
	var a = 5
	var b = 10
	var c = 15
	var d = 2
	var result = a + b - c * d

	fmt.Println(result)

	// augmented assignments
	var e = 10
	e += 10
	e += 5
	fmt.Println(e)

	// unary operator
	var i = 10
	i++
	fmt.Println(i)
}