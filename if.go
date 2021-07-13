package main

import "fmt"

func main() {

	name := " "

	if name == "Vien" {
		fmt.Println("Hello Vien")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hai, siapa namanya?")
	}

	if length := len(name); length > 5 {
		fmt.Println("its too long")
	} else {
		fmt.Println("OKE")
	}
}
