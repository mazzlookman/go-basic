package main

import "fmt"

func main() {
	registerUser("lukman", func(s string) bool {
		return s == "kasar"
	})
}

type BlackList func(string) bool

func registerUser(name string, bl BlackList) {
	if bl(name) {
		fmt.Println("You're blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}
