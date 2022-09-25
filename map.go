package main

import "fmt"

func main() {
	cars := map[string]string{
		"BMW":    "BMW M4 Competition",
		"Mercy":  "CLA 45 AMG",
		"Toyota": "Land Cruiser",
	}
	fmt.Println(cars)
	fmt.Println(cars["BMW"])
	fmt.Println(cars["Mercy"])

	book := make(map[string]string)
	book["title"] = "Go-Lang Dev"
	book["author"] = "Aqibmoh"
	book["wrong"] = "Ups"

	delete(book, "wrong")
	fmt.Println(book)
}
