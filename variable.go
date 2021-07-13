package main

import "fmt"

func main() {

	// harus menuliskan data type :
	var name string

	name = "Novien Ghoziana Indanartha"

	fmt.Println(name)

	// akan error kalau variable name diisi integer
	// name = 100

	// bisa mmebuat variable tanpa tipe data kalau diinit langsung nilainya :

	var age = 24

	fmt.Println(age)

	// di golang juga bisa membuat variable tanpa menuliskan "var"

	country := "Indonesia"

	fmt.Println(country)

	// Multiple Variable

	var (
		firstname = "Novien Ghoziana"
		lastname  = "Indanartha"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)

}
