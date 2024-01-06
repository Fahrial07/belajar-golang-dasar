package main

import "fmt"

func main() {

	name := "Dia Siapa ?"

	switch name {
	case "Eko":
		fmt.Println("Halo Eko")
	case "Budi":
		fmt.Println("Halo Budi")
	default:
		fmt.Println("Hi, Boleh Kenalan")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Terlalu Panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalau panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}
