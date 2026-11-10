package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("id is required")
	notFoundError = errors.New("id not found")
)

func getById(id string) error {
	if id == "" {
		return validationError
	}

	if id != "123" {
		return notFoundError
	}

	return nil
}

func main() {
	err := getById("")
	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Printf("validation error: %s", err.Error())
			} else if errors.Is(err, notFoundError) {
				fmt.Printf("not found error: %s", err.Error())
			} else {
				fmt.Println("Unknown error")
			}
	} else {
		fmt.Println("success")
	}
}