package main

import "fmt"

func main() {

	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var f = 3

	var c = a + b - d*e
	var g = c % f

	fmt.Println(g)

	var i = 10
	i += 10

	fmt.Println(i)

	i += 5

	fmt.Println(i)

	var j = 1
	j++ //j+1
	fmt.Println(j)

	j++

	fmt.Println(j)

	j-- // j-1

	fmt.Println(j)
}
