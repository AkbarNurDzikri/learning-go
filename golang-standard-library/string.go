package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Seafood", "foo"))
	fmt.Println(strings.Split("Dzikri Nur Akbar", " "))
	fmt.Println(strings.ToLower("DZIKRI NUR AKBAR"))
	fmt.Println(strings.ToUpper("dzikri nur akbar"))
	fmt.Println(strings.Trim(" Khalid Shalahuddin Akbar     ", " "))
	fmt.Println(strings.ReplaceAll("Dzikri Nur Akbar", "Dzikri Nur", "Bilal Shalahuddin"))
}