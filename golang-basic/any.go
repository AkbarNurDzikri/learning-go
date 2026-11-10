package main

import "fmt"

func ups() any {
	return "Anything"
}

func main() {
	var ups any = ups()
	fmt.Println(ups)
}