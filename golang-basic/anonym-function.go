package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) string {
	if blacklist(name) {
		return fmt.Sprintf("You're blocked, %v", name)
	} else {
		return fmt.Sprintf("Welcome, %v", name)
	}
}

func main() {
	// anonymous function bisa assign dalam variable
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	fmt.Println(registerUser("Anjing", blacklist))
	// anonymous function bisa langsung ditulis sebagai parameter
	fmt.Println(registerUser("Budi", func(name string) bool {
		return name == "Anjing"
	}))
}