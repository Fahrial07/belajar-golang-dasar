package main

import "fmt"

//operator asterisk atau *

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{"Kudus", "Jawa Tengah", "Indonesia"}
	address2 := &address1

	address2.City = "Semarang"

	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1) //ikut berbah
	fmt.Println(address2)
}
