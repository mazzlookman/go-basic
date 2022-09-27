package database

import "fmt"

var conn string

func init() {
	fmt.Println("Init dipanggil")
	conn = "MySQL"
}

func GetConnection() string {
	return conn
}
