package main

import "fmt"

type Hasname interface {
	GetName() string
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayPerson(value Hasname) string {
	return fmt.Sprintf("Hello, %v", value.GetName())
}

func main() {
	person := Person{Name: "Dzikri"}
	fmt.Println(SayPerson(person))
	
	animal := Animal{Name: "Kucing"}
	fmt.Println(SayPerson(animal))
}