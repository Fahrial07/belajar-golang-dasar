package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) {
	if blackList(name) {
		fmt.Println("You are black list", name)
	} else {
		fmt.Println("Welcone", name)
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Ali", blackList)

	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})

}
