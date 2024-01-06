package main

import "fmt"

type Customer struct {
	Name, Address, Email string
	Age                  int
	IsActive             bool
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Halo, ", name, "My Name is", customer.Name)
}

func main() {
	var ali Customer
	ali.Name = "Ali Fahrial Anwar"
	ali.Address = "Pati"
	ali.Age = 22
	ali.IsActive = true

	fmt.Println(ali)
	fmt.Println(ali.Name)
	fmt.Println(ali.Address)
	fmt.Println(ali.Age)
	fmt.Println(ali.IsActive)

	budi := Customer{
		Name:     "Budi Gunawan",
		Address:  "Balai Kota",
		Age:      40,
		IsActive: false,
	}

	joko := Customer{
		Name:     "Joko SI",
		Address:  "Balai Kota",
		Age:      40,
		IsActive: false,
	}

	fmt.Println(budi)

	budi.sayHello("Andi")
	joko.sayHello("Budi")
	ali.sayHello("Andi")

}
