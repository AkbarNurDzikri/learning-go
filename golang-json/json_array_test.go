package golangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "Udin",
		MiddleName: "Dindun",
		LastName: "Simalakadun",
		Hobbies: []string{"Game", "Reading", "Coding"},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{
		"FirstName": "Udin",
		"MiddleName": "Dindun",
		"LastName": "Simalakadun",
		"Age": 32,
		"Married": true,
		"Hobbies": ["Gaming", "Reading", "Coding"]
	}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Udin",
		MiddleName: "Dindun",
		LastName: "Simalakadun",
		Addresses: []Address{
			{
				Street: "Jl. Yuk No. 18",
				Country: "Karawang",
				PostCode: "41371",
			},
			{
				Street: "Jl. Bandar Juanda",
				Country: "Jakarta",
				PostCode: "430938",
			},
		},
	}

	bytes, err := json.Marshal(customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{
		"FirstName": "Udin",
		"MiddleName": "Dindun",
		"LastName": "Simalakadun",
		"Age": 32,
		"Married": true,
		"Hobbies": ["Gaming", "Reading", "Coding"]
	}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[
		{
			"Street": "Jl. Yuk No. 18",
			"Country": "Karawang",
			"PostCode": "41371"
		},
		{
			"Street": "Jl. Bandar Juanda",
			"Country": "Jakarta",
			"PostCode": "430938"
		}
	]`

	jsonBytes := []byte(jsonString)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(addresses)
}

func TestOnlyJsonArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street: "Jl. Yuk No. 18",
			Country: "Karawang",
			PostCode: "41371",
		},
		{
			Street: "Jl. Bandar Juanda",
			Country: "Jakarta",
			PostCode: "430938",
		},
	}

	bytes, err := json.Marshal(addresses)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}