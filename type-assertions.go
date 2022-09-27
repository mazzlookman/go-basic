package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	ta := random()
	switch ta.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	default:
		fmt.Println("Unknown")
	}
}
