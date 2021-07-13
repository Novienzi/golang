package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		// apapun yang dibawah break tidak akan dieksekusi
		fmt.Println("pengulangan ke", i)
	}
}
