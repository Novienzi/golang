package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func sayHello(customer Customer, name string) {
	fmt.Println("hello", name, "my name is", customer.Name)
}

func (customer Customer) greetingHello(name string) {
	fmt.Println("hello", name, "my name is", customer.Name)
}

func (person Customer) sayHmm() {
	fmt.Println("Hmmm from", person.Name)
}

func main() {
	var person1 Customer
	person1.Name = "vien"
	person1.Address = "Tulungung"
	person1.age = 24

	fmt.Println(person1)

	person2 := Customer{

		Name:    "Budi",
		Address: "Jakarta",
		age:     45,
	}

	fmt.Println(person2)

	// more sensitive bcs depends on struct decalaration
	person3 := Customer{"Joko", "Bandung", 54}

	fmt.Println(person3)

	sayHello(person1, "ani")

	person1.greetingHello("ami")

	person1.sayHmm()
}
