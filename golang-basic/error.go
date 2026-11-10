package main

import (
	"errors"
	"fmt"
)

func pembagian(nilai int, pembagi int) (int, error) {
	if(pembagi == 0) {
		return 0, errors.New("Pembagi 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	result, err := pembagian(5, 5)
	if(err == nil) {
		fmt.Println(fmt.Sprintf("Hasil pembagian, %v", result))
	} else {
		fmt.Println(fmt.Sprintf("Error: %v", err))
	}
}