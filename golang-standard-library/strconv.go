package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := strconv.ParseBool("salah")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

	// parse string to integer
	resultInt, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resultInt)
	}

	// parse integer to string
	strInt := strconv.Itoa(90000)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(strInt)
	}

	// contoh memformat numeric jadi binary
	binary := strconv.FormatInt(999,2)
	fmt.Println(binary)
}