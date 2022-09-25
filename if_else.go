package main

import "fmt"

func main() {
	name := "Mohammad"
	if name == "Mohammad" {
		fmt.Println("Hello Mohammad")
	} else if name == "Lukman" {
		fmt.Println("Hello Lukman")
	} else if name == "Aqib" {
		fmt.Println("Hello Aqib")
	} else {
		fmt.Println("Mohammad Lukman Aqib")
	}

	//if-else short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama kepanjangan")
	} else {
		fmt.Println("Pas lah")
	}
}
