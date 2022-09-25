package main

import "fmt"

func main() {
	//to make alias for a data type
	type NoHp int64
	type Active bool

	var nohp NoHp = 6281234567890
	var actv Active = true

	fmt.Println(nohp)
	fmt.Println(actv)

}
