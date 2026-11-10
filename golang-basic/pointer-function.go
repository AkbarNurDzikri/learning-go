package main

import "fmt"

type PointerAddress struct {
	City, Province, Country string
}

func changeCountryToIndonesia(address *PointerAddress) {
	address.Country = "Indonesia"
}

func main() {
	var address *PointerAddress = &PointerAddress{}
	changeCountryToIndonesia(address)

	fmt.Println(address)
}