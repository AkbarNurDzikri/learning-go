package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"Rudi", "Agus", "Ramdani", "Ujang", "Abud", "Roni"}
	values := []int{50,40,650,23,54,668,65}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Contains(names, "Udin"))
	fmt.Println(slices.Index(names, "Abud"))
	fmt.Println(slices.Index(names, "Obay"))
}