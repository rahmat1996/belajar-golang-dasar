package main

import (
	"fmt"
)

func main() {
	name := "Rahmat"
	country := "Indonesia"
	myName := func() {
		name = "Budi"
		country := "USA"
		fmt.Println(name)
		fmt.Println(country)
	}

	myName()

	fmt.Println(name)    // name will be changed from value on function
	fmt.Println(country) // country will not changed from value on function
}
