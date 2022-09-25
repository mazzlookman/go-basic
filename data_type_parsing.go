package main

import "fmt"

func main() {
	//number
	var values32 int32 = 280201
	var values64 int64 = int64(values32)
	var values16 int16 = int16(values32)

	fmt.Println(values32)
	fmt.Println(values64)
	fmt.Println(values16)

	//string
	name := "Hi Bell"
	e := name[0]
	eStr := string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eStr)
}
