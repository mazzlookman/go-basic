package main

import "fmt"

type Populace struct {
	Name    string
	Country string
}

func changePopulaceToIndo(populace *Populace) {
	populace.Country = "Indonesia"
}

func main() {
	//* penunjuk nilai
	//& penunjuk alamat
	//kalo mau nemuin *(nilai) harus ada &(alamat) nya
	p := Populace{"Mohammad", "USA"}
	changePopulaceToIndo(&p)
	fmt.Println(p)
}
