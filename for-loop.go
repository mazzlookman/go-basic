package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	//with statement
	for i := 1; i <= 10; i++ {
		fmt.Println("Looping ke: ", i)
	}

	//for range
	names := []string{"Mohammad", "Lukman", "Aqib"}
	for i, name := range names {
		fmt.Println("Index: ", i, "Nama: ", name)
	}
}
