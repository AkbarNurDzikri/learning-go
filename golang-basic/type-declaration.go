package main

import "fmt"

func main() {
	type NoKTP string

	var ktpDzikri NoKTP = "111111111111111"
	var contoh string = "222222222222222"

	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpDzikri)
	fmt.Println(contohKtp)

}