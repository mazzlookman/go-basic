package main

import "fmt"

func main() {
	runningApp(true)
	fmt.Println("Mohammad")
}

func endApp() {
	msg := recover()
	if msg != nil {
		fmt.Println("Error Message:", msg)
	}
	fmt.Println("End Program")
}

func runningApp(b bool) {
	defer endApp()
	if b {
		panic("Erorrr")
	}
	fmt.Println("App is running")
}
