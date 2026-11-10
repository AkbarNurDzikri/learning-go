package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	// cara menggunakan recover panic yang benar
	// dipanggil didalam defer function karena akan selalu dijalankan walau ada panic error
	message := recover()
	fmt.Println("Terjadi panic", message)
}

func runApp(error bool) {
	defer endApp() // key defer akan menjalankan function otomatis di akhir function ini
	if error {
		panic("Upsss Error cuy")
	}
	fmt.Println("Run App")

	// cara menggunakan recover panic yang salah
	// message := recover()
	// fmt.Println("Terjadi panic", message)
}

func main() {
	runApp(true)
	fmt.Println("Dzikri Nur Akbar")
}