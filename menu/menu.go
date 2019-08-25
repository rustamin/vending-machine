package main

import (
	"fmt"
)

func coba() {
	fmt.Println("test")
}

func menu() {
	fmt.Println("ini menu")
}

func mainMenu() {
	input := make([]int, 2)
	fmt.Print("Enter text: ")
	fmt.Scanln(&input[0], &input[1])

	if (input)[0] == 1 {
		fmt.Println("input nya 1")
		mainMenu()
	}

}

func main() {
	// input := make([]int, 2)
	// fmt.Print("Enter text: ")
	// fmt.Scanln(&input[0], &input[1])

	mainMenu()

	// switcher(&input)

	// fmt.Println(input)
	// fmt.Print("Enter text: ")
	// var input string
	// fmt.Scanln(&input)
	// fmt.Println(input)

	// fmt.Print("Enter text 2: ")
	// fmt.Scanln(&input)
	// fmt.Println(input)
}
