package main

import "fmt"

func getGoodBye(name string) string {
	return "Goodbye, " + name
}

func main() {
	// function bisa dijadikan sebagai value dari variable
	goodbye := getGoodBye
	fmt.Println(goodbye("CR7"))
}