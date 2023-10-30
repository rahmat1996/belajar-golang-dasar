package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak bisa 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {

	nilai, error := Pembagi(10, 0)
	if error == nil {
		fmt.Println("Hasil", nilai)
	} else {
		fmt.Println("Error", error.Error())
	}

}
