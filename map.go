package main

import "fmt"

func main() {
	// person := map[string]string = map[string]string

	// person["name"] = "Ali"
	// person["address"] = "Pati"

	person := map[string]string{
		"name":    "Ali Fahrial Anwar",
		"address": "Pati",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Ali"
	book["ops"] = "Salah"

	fmt.Println(book)

	delete(book, "ops")

	fmt.Println(book)

}
