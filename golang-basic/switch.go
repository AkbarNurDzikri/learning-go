package main

import "fmt"

func saySwitch(name string) string {
	switch name {
		case "Dzikri":
			return "Hi, Dzikri"
		case "Rudi":
			return "Hi, Rudi"
		default:
			return "Hi, what's your name?"
	}
}

// short statement switch
func switchShortStatement(name string) string {
	switch length := len(name); length > 5 {
	case true:
		return "Your name is too long"
	default:
		return "Your name is acceptable"
	}
}

// switch without condition
func switchWithoutCondition(name string) string {
	length := len(name)

	switch {
		case length > 10:
			return "Your name is too long"
		case length > 5:
			return "Your name is long"
		default:
			return "Your name is acceptable"
	}
}

func main() {
	fmt.Println(saySwitch("Dzikri"))
	fmt.Println(saySwitch("Rudi"))
	fmt.Println(saySwitch("Ucup"))

	fmt.Println(switchShortStatement("Dzikri"))
	fmt.Println(switchShortStatement("Rudi"))

	fmt.Println(switchWithoutCondition("Dzikri Nur Akbar"))
	fmt.Println(switchWithoutCondition("Rusdiana"))
	fmt.Println(switchWithoutCondition("Roy"))
}