package main

import "fmt"

func main() {
	// beberapa cara membuat map
	var person map[string]string = map[string]string{}
	person["name"] = "Dzikri"
	person["address"] = "Karawang"
	fmt.Println(person)

	var book = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Dzikri"
	book["publisher"] = "Erlangga"
	fmt.Println(book)

	// menghapus data pada map
	delete(book, "publisher")
	fmt.Println(book)
}