package main

import "fmt"

func getHello(name string) string {
	return "Hello" + name
}

func main() {
	result := getHello("vien")

	fmt.Println(result)
}
