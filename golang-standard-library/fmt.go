package main

import "fmt"

func main() {
	firstName := "Dzikri"
	lastName := "Nur Akbar"

	fmt.Println("Hi, '", firstName, lastName, "'")

	// substitusi print format
	fmt.Printf("Hi, '%s %s'", firstName, lastName)
}