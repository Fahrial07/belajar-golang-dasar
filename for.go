package main

import "fmt"

func main() {

	//for
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke ", counter)
	// 	counter++
	// }

	//init statement, post statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	fmt.Println("Selesai")

	//slice with for manual access
	names := []string{"Ali", "Fahrial", "Anwar"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for range
	//with index
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
	//not index
	for _, name := range names {
		fmt.Println(name)
	}
}
