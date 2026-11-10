package main

import "fmt"

// function bisa dijadikan sebagai parameter
func sayHelloWithFilter(name string, filter func(string) string) string {
	return "Hello, " + filter(name)
}

type Filter func(string) string

// daripada menuliskan data type yang panjang lebih baik di deklarasikann sebagai type diluar parameter
func sayHiWithFilter(name string, filter Filter) string {
	return "Hi, " + filter(name)
}


func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	filter := spamFilter
	fmt.Println(sayHelloWithFilter("Anjing", filter))

	fmt.Println(sayHiWithFilter("Rudi", spamFilter))
}