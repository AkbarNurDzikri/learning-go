package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Karawang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // pass by value

	// address1.City = "Bandung"

	// fmt.Println(address1) // data tidak berubah
	// fmt.Println(address2) // data city berubah karena pengiriman data dari variable adalah passing by value, bukan reference

	// pass by reference
	address3 := &address1 // data akan berubah karena menggunakan & sebagai reference
	address3.City = "Sumedang" // address1.City ikut berubah
	// fmt.Println(address1)
	// fmt.Println(address3)

	// deklarasi type data pointer menggunakan *
	var address4 *Address = &address1
	address4.City = "Bekasi"
	fmt.Println(address1)
	fmt.Println(address4)
	

}