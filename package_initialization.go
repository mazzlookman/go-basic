package main

import (
	"fmt"
	"golang-dasar/database"
)

func main() {
	connection := database.GetConnection()
	fmt.Println(connection)
}
