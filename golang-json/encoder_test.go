package golangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestEncoder(t *testing.T) {
	writer, err := os.Create("CustomerOut.json")
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName: "Udin",
		MiddleName: "Dindun",
		LastName: "Simalakadun",
	}

	encoder.Encode(customer)
}