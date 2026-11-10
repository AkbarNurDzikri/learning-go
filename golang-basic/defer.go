package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func RunApplication() {
	defer logging() // key defer akan menjalankan function otomatis di akhir function ini
	fmt.Println("Running application")
}

func main() {
	RunApplication()
}