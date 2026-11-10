package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Dzikri"
	names[1] = "Nur"
	names[2] = "Akbar"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array langsung
	var fixedValues = [3]int{
		90,
		80,
	}

	// menampilkan array, index ke 2 by default akan bernilai 0 (string kosong jika string) jika tidak diisi
	fmt.Println(fixedValues)

	fixedValues[2] = 70

	fmt.Println(fixedValues[0])
	fmt.Println(fixedValues[1])
	fmt.Println(fixedValues[2])

	// dynamic len values
	var dynamicValues = [...]int{90, 80, 70, 60, 50}
	fmt.Println(dynamicValues)

	// menampilkan panjang array
	fmt.Println("Panjang array : ", len(dynamicValues))
}