package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {

	// if the filter function have many parameter type you can use type declarations
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {

	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("vien", spamFilter)
	sayHelloWithFilter("anjing", spamFilter)
}
