package main

import "fmt"

type NewAddress struct {
	City, Province, Country string
}

func main() {
	// membuat data/object kosong dengan cara manual
	var address1 *NewAddress = &NewAddress{}
	var address2 *NewAddress = address1
	fmt.Println(address1)
	fmt.Println(address2)

	// membuat data/object kosong menggunakan new
	var address3 *NewAddress = new(NewAddress)
	var address4 *NewAddress = address3
	fmt.Println(address3)
	fmt.Println(address4)
}