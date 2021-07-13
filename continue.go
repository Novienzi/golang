package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		// apapun yang dibawah continue tidak dieksekusi  dan langsung menjalankan post statement
		fmt.Println("pengulangan ke", i)
	}
}
