package main

import "fmt"

func main() {
	//math operations
	a := 10
	b := 5
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)

	//augmented assigment
	// a+=b => a=a+b
	var c = 28
	c += a
	fmt.Println(c)

	//unary operator
	// ++, --, -, +, !
	d := 1
	d++
	d++
	fmt.Println(d)

	//comparison operator
	// < | > | == | <= | >= | !=
	name1 := "Hi Bell"
	name2 := "Hi Bell"

	result := name1 == name2
	fmt.Println(result)

	//boolean operator
	// && | || | !
	nilaiAkhir := 90
	absensi := 80

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := absensi >= 80

	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}
