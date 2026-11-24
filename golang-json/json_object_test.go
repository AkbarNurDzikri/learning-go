package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street string
	Country string
	PostCode string
}

type Customer struct {
	FirstName string
	MiddleName string
	LastName string
	Hobbies []string
	Addresses []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Udin",
		MiddleName: "Dindun",
		LastName: "Simalakadun",
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}