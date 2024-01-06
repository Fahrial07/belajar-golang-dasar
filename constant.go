package main

import "fmt"

func main() {
	const firstName string = "Ali"
	const middleName = "Fahrial"
	const lastName = "Anwar"

	// firstName = "OKE"

	fmt.Println(firstName, middleName)

	const (
		fruit string = "Anggur"
		price        = 25000
	)

	fmt.Println(fruit, price)
}
