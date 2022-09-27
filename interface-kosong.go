package main

import "fmt"

func EmptyIface(i int) interface{} {
	if i == 1 {
		return true
	} else if i == 2 {
		return 2
	} else {
		return "Ups"
	}
}

func main() {
	iface := EmptyIface(1)
	fmt.Println(iface)
}
