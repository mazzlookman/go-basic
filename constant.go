package main

import "fmt"

func main() {
	const firstName string = "Kepo"
	const lastName = "Kepo Lagi"

	fmt.Println(firstName)
	fmt.Println(lastName)

	//constant's value can't be changed
	//error
	//firstName = "changing"
	//lastName = "changing"

	//multiple constant
	const (
		drink string = "Hot Coffee"
		like         = true
	)
	fmt.Println(drink)
	fmt.Println(like)
}
