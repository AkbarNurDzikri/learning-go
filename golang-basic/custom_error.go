package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func updateData(id string, data any ) error {
	if id == "" {
		return &validationError{"id is required"}
	} else if id != "123" {
		return &notFoundError{"id not found"}
	} else {
		return nil
	}
}

func main() {
	err := updateData("", nil)
	if err != nil {
		// if validationError, ok := err.(*validationError); ok {
		// 	fmt.Println(fmt.Sprintf("Validation error: %v", validationError.Message))
		// 	} else if notFoundError, ok := err.(*notFoundError); ok {
		// 	fmt.Println(fmt.Sprintf("Not found error: %v", notFoundError.Message))
		// } else {
		// 	fmt.Println("Unknown error")
		// }

		// menggunakan switch type assertion
		switch finalError := err.(type) {
		case *validationError:
				fmt.Println(fmt.Sprintf("Validation error: %v", finalError.Message))
		case *notFoundError:
				fmt.Println(fmt.Sprintf("Not found error: %v", finalError.Message))
		default:
			fmt.Println("Unknown error")
		}
	} else {
		fmt.Println("success")
	}
}