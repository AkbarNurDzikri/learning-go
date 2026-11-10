package main

import "fmt"

type AsteriskAddress struct {
	City, Province, Country string
}

func main() {
	var addressX AsteriskAddress = AsteriskAddress{"Malang", "Jawa Timur", "Indonesia"}
	var addressY *AsteriskAddress = &addressX
	addressY.City = "Pemalang" // addressX akan berubah karena deklarasi addressY menggunaan pointer
	addressY.Province = "Jawa Tengah" // addressX akan berubah karena deklarasi addressY menggunaan pointer
	fmt.Println(addressX)
	fmt.Println(addressY)

	*addressY = AsteriskAddress{"Karawang", "Jawa Barat", "Indonesia"}
	fmt.Println("AddressX ", addressX)
	fmt.Println("AddressZ ", addressY)
}