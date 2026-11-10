package main

import "fmt"

func forBreak() {
	for i := 0; i < 10; i++ {
		// jika i sama dengan 5, perulangan dihentikan
		if i == 5 {
			break
		}

		fmt.Println(i)
	}
}

func forContinue() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)
	}
}

func main() {
	// forBreak()
	forContinue()
}