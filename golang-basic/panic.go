package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp() // key defer akan menjalankan function otomatis di akhir function ini
	if error {
		panic("Upsss Error cuy")
	}
	fmt.Println("Run App")
}

func main() {
	runApp(true)
}