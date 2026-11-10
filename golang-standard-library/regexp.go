package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile(`d[a-z]i`)

	fmt.Println(regex.MatchString("dzikri"))
	fmt.Println(regex.MatchString("doni"))
	fmt.Println(regex.MatchString("diki"))
	fmt.Println(regex.FindAllString("debi dobi dzikri doni dedi dandi dini d3ki", 10))
}