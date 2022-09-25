package main

import "fmt"

func main() {
	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daysSlice1 := days[5:]
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Lembur"
	daysSlice1[1] = "Minggu Libur"
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Senin Libur")
	fmt.Println(daysSlice1)

	daysSlice2[0] = "Ups"
	fmt.Println(daysSlice2)
	fmt.Println(days)

	fmt.Println("===================")

	//make new slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Mohammad"
	newSlice[1] = "Lukman"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fmt.Println("===================")

	//copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

	//array and slice is different
	iniArray := [...]int{1, 2, 3} //array eksplisit menyebutkan length nya
	iniSlice := []int{1, 2, 3}    // sedangkan slice tidak usah (fleksibel)
}
