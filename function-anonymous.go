package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blackList Blacklist) {
	if blackList(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {

	blackList := func(name string) bool {
		return name == "nono"
	}

	registerUser("vien", blackList)
	registerUser("nono", blackList)
	registerUser("admin", blackList)

	registerUser("nono", func(name string) bool {
		return name == "nono"
	})

}
