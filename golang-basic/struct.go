package main

import "fmt"

type Customer struct {
	Name, Address  string
	Age int
}

// field struct bisa juga berisi method (function)
// cara membuatnya seperti ini
func (customer Customer) SayHello(name string) string {
	return fmt.Sprintf("Hello, %s, My name is %s.", name, customer.Name)
}

func main() {
	var customer1 Customer

	// banyak cara untuk membuat data dari struct
	// salah satunya seperti ini
	customer1.Name = "Dzikri Nur Akbar"
	customer1.Address = "Indonesia"
	customer1.Age = 30

	fmt.Println(customer1)
	fmt.Println(customer1.Name)
	fmt.Println(customer1.Address)
	fmt.Println(customer1.Age)
	
	// bisa juga seperti ini
	var customer2 Customer = Customer{
		Name: "Udin Dindun",
		Address: "Surabaya",
		Age: 35,
	}

	fmt.Println(customer2)
	fmt.Println(customer2.Name)
	fmt.Println(customer2.Address)
	fmt.Println(customer2.Age)
	
	// bisa juga seperti ini dengan langsung menulis value tanpa fieldnya
	// dengan syarat harus berurutan
	customer3 := Customer{
		"Radi",
		"Bekasi",
		25,
	}
	
		fmt.Println(customer3)
		fmt.Println(customer3.Name)
		fmt.Println(customer3.Address)
		fmt.Println(customer3.Age)

		// memanggil method sayHello
		fmt.Println(customer3.SayHello("Rudi"))

}