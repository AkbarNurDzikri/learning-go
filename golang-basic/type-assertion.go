package main

import "fmt"

func random() interface{} {
	return true
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Println(resultString)

	// var resultInt int = result.(int)
	// fmt.Println(resultInt)

	// memeriksa tipe data menggunakan switch	
	switch value := result.(type) {
	case string:
		fmt.Println("string ", value)
	case int:
		fmt.Println("int ", value)
	default:
		fmt.Println("Unknown")
	}
}