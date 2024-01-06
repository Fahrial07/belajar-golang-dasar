package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("EKO")
	fmt.Println(result)

	fmt.Println(helper.Application)

	fmt.Println(helper.version)           //tidak bisa diakses
	fmt.Println(helper.sayGoodBye("YOU")) //tidak bisa di akses

}
