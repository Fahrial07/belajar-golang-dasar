package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Ali"
	names[1] = "Fahrial"
	names[2] = "Anwar"

	var values = [3]int{
		90,
		80,
		95,
	}

	numbers := [3]int{
		1,
		2,
		3,
	}

	dynamic := [...]int{
		90,
		80,
		70,
		100,
		20,
	}

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(numbers)

	fmt.Println(dynamic)

}
