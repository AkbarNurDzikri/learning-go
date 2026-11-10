package main

import (
	"container/list"
	"fmt"
)

func main() {
	var data *list.List = list.New()
	data.PushBack("Dzikri")
	data.PushBack("Nur")
	data.PushBack("Akbar")

	// var head *list.Element = data.Front()
	// fmt.Println(head.Value) // Dzikri

	// next := head.Next()
	// fmt.Println(next.Value) // Nur

	// next = next.Next()
	// fmt.Println(next.Value) // Akbar

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}