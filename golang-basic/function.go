package main

import "fmt"

func sayHi() {
	fmt.Println("Hi, bro!")
}

func sayHiTo(firstName string, lastName string) {
	fmt.Println("Hi, ", firstName, lastName)
}

func getHello(name string) string {
	hello := "Hello, " + name
	return hello
}

func getFullName(firstName string, lastName string) (string, string) {
	return firstName, lastName
}

func namedReturnValue() (firstName, midName, lastName string) {
	firstName = "Dzikri"
	midName = "Nur"
	lastName = "Akbar"
	return firstName, midName, lastName
}

func main() {
	sayHi()
	sayHiTo("Dzikri", "Akbar")
	fmt.Println(getHello("Alice"))

	// menngambil multiple return value
	firstName, lastName := getFullName("Dzikri", "Akbar")
	fmt.Println("Full Name: ", firstName, lastName)

	// mengambil salah satu return multiple value
	firstNameOnly, _ := getFullName("Dzikri", "Akbar")
	fmt.Println("First Name Only: ", firstNameOnly)

	a,b,c := namedReturnValue()
	fmt.Println(a,b,c)
}