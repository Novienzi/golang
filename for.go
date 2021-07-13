package main

import "fmt"

func main() {

	counter := 1

	for counter < 10 {
		fmt.Println("pengulangan ke ", counter)
		counter++
	}

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Ini pengulangan kedua ke", counter)
	}

	slice := []string{"ani", "budi", "cahya", "deny"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(i)
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index ke", i, "=", value)
		fmt.Println(value)
	}

	names := map[string]string{
		"name":       "zia",
		"background": "Engineer",
	}

	for key, value := range names {
		fmt.Println(key, "=", value)
	}

	for _, value := range names {
		fmt.Println(value)
	}
}
