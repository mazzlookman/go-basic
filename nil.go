package main

import "fmt"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	m := newMap("")
	if m == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(m)
	}
}
