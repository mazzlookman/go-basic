package main

import (
	"fmt"
)

func main() {
	var aqib Customer
	aqib.Name = "Mohammad Lukman"
	aqib.Address = "Pemalang"
	aqib.Age = 21
	fmt.Println(aqib)

	bella := Customer{
		Name:    "Bella",
		Address: "Pemalang",
		Age:     21,
	}
	fmt.Println(bella.Name)

	//ini rentan error, data harus berurutan
	kumis := Customer{"Kumis", "Kidul", 22}
	fmt.Println(kumis)

	aqib.checkout("Joko")
}

type Customer struct {
	Name, Address string
	Age           int
}

func (c Customer) checkout(name string) {
	fmt.Println(name, "Checkoutin", c.Name)
}
