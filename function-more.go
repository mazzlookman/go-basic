package main

import "fmt"

func main() {
	sum := sumAll(1, 2, 3, 4, 5, 6, 10)
	fmt.Println(sum)

	//slice parameter
	numbers := []int{1, 2, 3, 4, 5, 6, 10}
	total := sumAll(numbers...)
	fmt.Println(total)

	//function value
	gb := sayGoodBye
	fmt.Println(gb("Mohammad"))
}

//variadic function
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//function value
func sayGoodBye(name string) string {
	return "Good Bye " + name
}
