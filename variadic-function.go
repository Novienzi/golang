package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	total := sumAll(10, 12, 13)

	fmt.Println(total)

	slice := []int{1, 2, 4, 6}

	total = sumAll(slice...)
	fmt.Println(total)
}
