package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	helper.HaloGans("Mohammad")
	//helper.haloGans("name") //error

	fmt.Println(helper.App)
	//fmt.Println(helper.version)//error
}
