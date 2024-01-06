package main

import "fmt"

func main() {
	names := [...]string{"Ali", "Fahrial", "Anwar", "Abid", "Cello", "Melvin"}

	slice1 := names[4:6]

	fmt.Println(slice1)

	slice2 := names[:3]

	fmt.Println(slice2)

	slice3 := names[3:]

	fmt.Println(slice3)

	slice4 := names[:]

	fmt.Println(slice4)

	var manualSlice []string = names[:]

	fmt.Println(manualSlice)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:] // Sabtu, Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Libur"
	daySlice1[1] = "Minggu Libur"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice2)
	fmt.Println(daySlice1)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)

	newSlice[0] = "Ali"
	newSlice[1] = "Ali"
	//newSlice[2] = "Ali" ini akan error, karena telah di tentukan capacity slice 2, dan harusnya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Ali")

	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Budi"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	//copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
