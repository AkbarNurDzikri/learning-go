package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name": name,
		}
	}
}

func main() {
	data := newMap("Dzikri")
	if data == nil {
		fmt.Println("Data masih kosong")
	} else {
		fmt.Println(data["name"])
	}
}