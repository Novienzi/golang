package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "novien",
		"address": "Tulungagung",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	fmt.Println(len(person))

	newPerson := make(map[string]string)
	newPerson["name"] = "Riza"
	newPerson["address"] = "Bandung"
	newPerson["noHape"] = "082232645458"

	fmt.Println(newPerson)

	delete(newPerson, "noHape")

	fmt.Println(newPerson)

}
