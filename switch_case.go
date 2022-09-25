package main

import "fmt"

func main() {
	name := "Mohammad"

	switch name {
	case "Mohammad":
		fmt.Println("Hello Mohammad")
	case "Lukman":
		fmt.Println("Hello Lukman")
	case "Aqib":
		fmt.Println("Hello Aqib")
	default:
		fmt.Println("Mohammad Lukman Aqib")
	}

	//switch tanpa kondisi
	switch {
	case name == "Mohammad":
		fmt.Println("Hello Mohammad")
	case name == "Lukman":
		fmt.Println("Hello Lukman")
	case name == "Aqib":
		fmt.Println("Hello Aqib")
	default:
		fmt.Println("Mohammad Lukman Aqib")
	}

	//short switch-case
	switch length := len(name); length > 5 {
	case true:
		fmt.Println(">5")
	case false:
		fmt.Println("Pas")
	}
}
