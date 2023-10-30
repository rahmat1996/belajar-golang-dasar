package main

import (
	"fmt"
)

func main() {
	var name string = "Rahmat"

	switch name {
	case "Rahmat":
		fmt.Println("Hi Rahmat")
	case "Budi":
		fmt.Println("Hi Budi")
	default:
		fmt.Println("Namamu belum ada!")
		fmt.Println("Kenalan donk!")
	}

	//short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Name too long!")
	case false:
		fmt.Println("Good!")
	}

	//switch without condition
	value := 91
	switch {
	case value > 90:
		fmt.Println("Excelent!")
	case value >= 70:
		fmt.Println("Good.")
	case value >= 0 && value < 70:
		fmt.Println("Bad")
	default:
		fmt.Println("Something wrong on value")
	}

}
