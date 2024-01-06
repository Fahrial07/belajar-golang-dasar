package main

import "fmt"

func logging() {
	fmt.Println("Selesai meanggil function logging")
}

func runApplication() {
	defer logging()

	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
