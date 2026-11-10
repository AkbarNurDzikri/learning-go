package main

import "fmt"

// parameter varArgs bisa menerima lebih dari 1 parameter
func sumAll(numbers ...int) int {
	total := 0

for _, number := range numbers {
	total += number
}

	return total
}

func main() {
	// bisa mengirim lebih dari 1 parameter ke function yang menerima parameter varArgs
	// tanpa harus mengubah parameter sebagai array atau slice
	result := sumAll(10,10,10)
	fmt.Println(result)

	// jika ternyata sudah terlanjut punya data berupa array atau slice
	// gunakan titik tiga setelah nama parameter
	var numbers []int = []int{1,2,3,4,5}
	fmt.Println(sumAll(numbers...))
}