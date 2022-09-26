package main

import "fmt"

func main() {
	sayHello()
	sayHi("Mohammad")

	result := getHello("Lukman")
	fmt.Println(result)

	fn, ln := getFullName()
	fmt.Println(fn, ln)

	fn2, _ := getFullName()
	fmt.Println(fn2)

	first, middle, last := getCompleteName()
	fmt.Println(first, middle, last)
}

func sayHello() {
	fmt.Println("Hello Gan")
}

func sayHi(name string) {
	fmt.Println("Hai ", name)
}

func getHello(name string) string {
	return "Hello " + name
}

func getFullName() (string, string) {
	return "Mohammad ", "Lukman"
}

func getCompleteName() (fn, mn, ln string) {
	fn = "Mohammad"
	mn = "Lukman"
	ln = "Aqib"

	return fn, mn, ln
}
