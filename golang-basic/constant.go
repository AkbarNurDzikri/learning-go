package main

import "fmt"

func main() {
	// constant jika tidak dipakai tidak menyebabkan error
	// constant tidak bisa diubah nilainya
	const firstName = "Dzikri Nur Akbar"
	const age = 30

	// error constant tidak boleh diubah
	// firstName = "Khalid Shalahuddin Akbar"
	// age = 25

	// sama seperti variable, constant juga bisa declare multiple variables
	const (
		address = "Jl. Raya Cibubur"
		phoneNumber = "081234567890"
	)

	fmt.Println(address)
	fmt.Println(phoneNumber)
}