package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "Novien"
	names[1] = "Ghoziana"
	names[2] = "Indanartha"

	fmt.Println(names)

	values := [3]int{
		100,
		99,
		76,
	}

	fmt.Println(values)

	values[0] = 85

	fmt.Println(values)
	fmt.Println(len(names))
	fmt.Println(names[2])

}
