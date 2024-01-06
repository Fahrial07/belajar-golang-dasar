package main

import "fmt"

func main() {
	name := "Raden"

	if name == "Ali" {
		fmt.Println("Hello", name)
	} else if name == "Joko" {
		fmt.Println("Hello", name)
	} else if name == "Andi" {
		fmt.Println("Hello", name)
	} else {
		fmt.Println("Hi, Boleh kanalan ?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
