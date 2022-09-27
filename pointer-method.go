package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	m := Man{"Mohammad"}
	m.Married()
	fmt.Println(m)
}
