package main

import "fmt"

func endApp() {

	message := recover()

	if message != nil {
		fmt.Println("error with", message)
	}

	fmt.Println("app is finished")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("error")
	}

	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
}
