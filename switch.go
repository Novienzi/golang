package main

import "fmt"

func main() {

	name := "vien"

	switch name {
	case "vien":
		fmt.Println("HAI VIEN!")
	case "budi":
		fmt.Println("HAI budi!")
	default:
		fmt.Println("hai, kamu siapa namanya?")
	}

	switch length := len(name); length > 10 {
	case true:
		fmt.Println("Its too long")
	case false:
		fmt.Println("OKE")
	}

	long := len(name)

	switch {
	case long > 10:
		fmt.Println("panjang banget!")
	case long > 5:
		fmt.Println("masih panjangg :(")
	default:
		fmt.Println("oke")

	}

}
