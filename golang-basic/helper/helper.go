package helper

import "fmt"

var version = "1.0.0"
var Application = "Golang"

// tidak bisa diakses dari luar package helper
func sayGoodBye(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

// bisa diakses dari luar package helper
func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}