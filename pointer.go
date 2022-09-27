package main

import "fmt"

type Address struct {
	Name string
	Addr string
}

func main() {
	addr1 := Address{"Mohammad", "Pemalang"}
	addr2 := &addr1

	addr2.Addr = "Purwakarta"

	fmt.Println(addr1) //berubah
	fmt.Println(addr2)

	//contoh 2
	ala1 := Address{"Bella", "Pemalang"}
	ala2 := &ala1
	ala2.Addr = "Bekasi"

	ala2 = &Address{"Pram", "Pemalang"}

	fmt.Println(ala1) // tidak berubah
	fmt.Println(ala2) //ini berubah

	//contoh 3
	adr1 := Address{"Lukman", "Jakarta"}
	adr2 := &adr1
	adr2.Addr = "Singapore"

	*adr2 = Address{"Aqib", "USA"}

	fmt.Println(adr1)
	fmt.Println(adr2)

	//contoh 4 (keyword new())
	alamat := new(Address)
	fmt.Println(alamat)
}
