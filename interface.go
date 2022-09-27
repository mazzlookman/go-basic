package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayName(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (p Person) GetName() string {
	return p.Name
}

type Animal struct {
	Name string
}

func (a Animal) GetName() string {
	return a.Name
}

func main() {
	aqib := Person{Name: "Aqib"}
	sayName(aqib)

	cat := Animal{Name: "Pushh"}
	sayName(cat)
}
