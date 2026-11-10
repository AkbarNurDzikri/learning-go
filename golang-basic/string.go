package main

import "fmt"

func main() {
	fmt.Println("Dzikri")
	fmt.Println("Nur")
	fmt.Println("Akbar")

	fmt.Println(len("Dzikri"))
	fmt.Println("Nur"[0]) // saat mengambil index string, hasilnya adalah byte, bukan string
	fmt.Println("Akbar"[1]) // saat mengambil index string, hasilnya adalah byte, bukan string
}