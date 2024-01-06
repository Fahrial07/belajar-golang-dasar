package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Ali"
	middleName = "Fahrial"
	lastName = "Anwar"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompleteName()

	fmt.Println(a, b, c)
}
