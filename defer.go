package main

import "fmt"

func main() {
	runApp(0)
}

func logging() {
	fmt.Println("Program End")
}

func runApp(value int) {
	defer logging()
	fmt.Println("Program is running")
	result := 10 / value
	fmt.Println(result)
}
