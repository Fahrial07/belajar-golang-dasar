package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Pati", "Jawa Tengah", "Indonesia"}
	// address2 := address1

	// address2.City = "Demak"

	// fmt.Println(address1) //value tidak berubah
	// fmt.Println(address2) //baru berubah city menjadi Demak

	address1 := Address{"Pati", "Jawa Tengah", "Indonesia"}
	address2 := &address1 //pointer / reference ke address1 dan tidak duplikat lagi

	//manual / biasa
	//var address1 Address = Address{"Pati", "Jawa Tengah", "Indonesia"}
	//var address2 *Address = &address1 //pointer / reference ke address1 dan tidak duplikat lagi

	address2.City = "Demak"

	fmt.Println(address1) //value  berubah
	fmt.Println(address2) //baru berubah city menjadi Demak
}
