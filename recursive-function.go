package main

import "fmt"

func main() {
	fk := factorialRecursive(5)
	fmt.Println(fk)
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}
