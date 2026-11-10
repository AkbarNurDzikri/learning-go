package main

import "fmt"

func forLoop(counter int, max int) {
	fmt.Println("Start For Loop")

	for counter <= max {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("Finish For Loop")
}

func forLoopInitStatement(counter int, max int) {
	fmt.Println("Start with init statement")

	for i := counter; i <= max; {
		fmt.Println(i)
		i++
	}

	fmt.Println("Finish with init statement")
}

func forLoopInitPostStatement(counter int, max int) {
	fmt.Println("Start with init and post statement")

	for i := counter; i <= max; i++ {
		fmt.Println(i)
	}

	fmt.Println("Finish with init and post statement")
}

func forRange() {
	var names []string = []string{"Dzikri", "Nur", "Akbar"}

	// manual
	// for i := 0; i < len(names); i++ {
	// 	fmt.Println(names[i])
	// }

	// for range
	// use index (or key if map) and value for access data
	for i, name :=  range names {
		fmt.Println("index: ", i, " | value: ", name)
	}

	// index unused
	for _, name := range names {
		fmt.Println("value: ", name)
	}

	var fruits map[string]string = make(map[string]string)
	fruits["apple"] = "red"
	fruits["banana"] = "yellow"
	fruits["orange"] = "orange"

	// for range from map
	for key, value := range fruits {
		fmt.Println("key: ", key, " | value: ", value)
	}
}

func main() {
	// forLoop(1, 10)
	// forLoopInitStatement(1, 20)
	// forLoopInitPostStatement(1, 30)
	forRange()
}