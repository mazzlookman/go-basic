package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Can't divide zero ")
	} else {
		r := a / b
		return r, nil
	}
}

func main() {
	hasil, err := divide(2, 0)
	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println("Errorrr", err.Error())
	}
}
