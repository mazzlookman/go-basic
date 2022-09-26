package main

import "fmt"

func main() {
	filterBadWord("Kasar", wordFilter)
}

//use type declarations
type Filter func(string) string

//func filterBadWord(name string, filter func(name string) string) {
//	ft := filter(name)
//	fmt.Println("Hello", ft)
//}

func filterBadWord(name string, filter Filter) {
	ft := filter(name)
	fmt.Println("Hello", ft)
}

func wordFilter(name string) string {
	if name == "Kasar" {
		return "Ih kasar yah"
	} else {
		return name
	}
}
