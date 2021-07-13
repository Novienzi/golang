package main

import "fmt"

func main() {

	days := [...]string{
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}

	slice := days[2:7]

	fmt.Println(slice)
	fmt.Println(days)
	fmt.Println(slice[0])

	// will affect to data of days either
	// slice[1] = "KAMIS LAGI"

	fmt.Println(days) // [senin selasa rabu KAMIS LAGI jumat sabtu minggu]
	fmt.Println(slice[1])
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	// different with append, it will add new data if the capacity of slice is full.
	daySlice2 := append(slice, "After Minggu")

	fmt.Println(daySlice2)

	// if there any update, it will not affect the data of days
	daySlice2[2] = "test aja si"

	fmt.Println(daySlice2)

	newSlice := make([]string, 3, 8)

	newSlice[0] = "Saya"
	newSlice[1] = "ITU"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	copy(toSlice, fromSlice)

	fmt.Print(toSlice)

	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println("ini", iniArray)
	fmt.Println(iniSlice)
}
