package main

import "fmt"

//any = interface{} / interface kosong
func Ups() any {
	//return 1
	return true
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}
