package main

import "fmt"

func main() {
	type NoKTP string

	var ktpAli NoKTP = "12321312312321"
	var contohKTP NoKTP = NoKTP("1121312321")

	fmt.Println(ktpAli)
	fmt.Println(contohKTP)

}
