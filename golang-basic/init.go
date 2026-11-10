package main

import (
	"belajar-golang-dasar/database"
	_ "belajar-golang-dasar/internal" // pemanggilan package tanpa method (yang otomatis menjalankan init)
	"fmt"
)

func main() {
	fmt.Println(database.GetDatabase())
}