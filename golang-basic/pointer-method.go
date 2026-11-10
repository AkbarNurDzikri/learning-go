package main

import "fmt"

type Man struct {
	name string
}

// gunakan selalu pointer di parameter method agar mudah
func (man *Man) married() {
	man.name =  fmt.Sprintf("Mr. %s", man.name)
}

func main () {
	dzikri := Man{"Dzikri"}
	dzikri.married()

	fmt.Println(dzikri.name)
}