package main

import "fmt"

func endApp() {
	fmt.Println("End App")

	message := recover()
	fmt.Println("Error with message:", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("UPS ERROR")
	}

	// message := recover()
	// fmt.Println("Error with message:", message)
}

func main() {
	// runApp(false)
	runApp(true)
	fmt.Println("Ali Fahrial Anwar")
}
