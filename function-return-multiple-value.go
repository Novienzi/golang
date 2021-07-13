package main

import "fmt"

func getFullName() (string, string) {
	return "novien", "ghoziana"
}

func main() {
	// firstName, lastName := getFullName()
	firstName, _ := getFullName()
	// fmt.Println(firstName, lastName)
	fmt.Println(firstName)

}
