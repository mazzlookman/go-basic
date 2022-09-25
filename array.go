package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Mohammad"
	names[1] = "Lukman"
	names[2] = "Aqib"

	fmt.Println(names[0])

	//deklarasi langsung
	var values = [3]int{
		80, 90, 95,
	}
	fmt.Println(values[1])

	//[...]
	var city = [...]string{
		"Pemalang", "Comal", "Tegal",
	}
	fmt.Println(city)
}
