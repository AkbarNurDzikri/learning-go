package main

import "fmt"

func main() {
	var name string

	name = "Dzikri Nur Akbar"
	fmt.Println(name)

	name = "Khalid Shalahuddin Akbar"
	fmt.Println(name)

	// jika variable langsung di assign data, boleh tidak perlu declare tipe data
	var umur = 30
	fmt.Println(umur)

	// saat membuat variable, keyword var tidak wajib digunakan, namun kita perlu menulis := sebagai gantinya (untuk first time declaration)
	// untuk mengganti nilainya, gunakan =
	jenisKelamin := "Laki-laki"
	fmt.Println(jenisKelamin)

	// multi variables declaration
	var (
		firstName = "Dzikri"
		lastName  = "Nur Akbar"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}