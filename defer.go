package main

import "fmt"

func logging() {
	fmt.Println("done call function")
}

func runApp(value int) {
	defer logging()
	fmt.Println("run application")
	result := 10 / value

	fmt.Print(result)

}

func main() {
	runApp(0)
}
