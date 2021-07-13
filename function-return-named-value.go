package main

import "fmt"

func getDataPerson() (no int, fullname string, address string) {
	no = 001
	fullname = "Novien Ghoziana I"
	address = "Tulungagung"

	return
}

func main() {

	fmt.Println(getDataPerson())

}
