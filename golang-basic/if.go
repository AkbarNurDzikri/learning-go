package main

import "fmt"

func sayHello(name string) string {
	if name == "Dzikri" {
		return "Hello Dzikri"
	} else if name == "Bilal" {
		return "Hello Bilal"
	} else {
		return "Hello, what's your name?"
	}
}

func shortStatement(name string) string {
	if length := len(name); length > 5 {
		return fmt.Sprintf("Hi, %s, your name is too long", name)
	}

	return fmt.Sprintf("Hi %s, your name is acceptable", name)
}


func main() {
	fmt.Println(sayHello("Dzikri"))
	fmt.Println(sayHello("Bilal"))
	fmt.Println(sayHello("Khalid"))

	fmt.Println(shortStatement("Dzikri"))
	fmt.Println(shortStatement("Kiki"))
}