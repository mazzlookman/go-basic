package main

import "fmt"

func main() {
	//Variable declaration
	//1. var with data type declaration
	var name string
	name = "Phi Cool"
	fmt.Println(name)

	name = "Suha Day"
	fmt.Println(name)

	//2. only var
	var age = 20
	fmt.Println(age)

	//3. only use :=
	address := "wakanda forevah"
	fmt.Println(address)
	girlFriend := false
	fmt.Println(girlFriend)

	//4. multiple variable
	var (
		drink = "Coffe"
		hot   = true
	)
	fmt.Println(drink)
	fmt.Println(hot)
}
