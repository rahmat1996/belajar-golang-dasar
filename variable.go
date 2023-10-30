package main

import "fmt"

func main() {
	var name string

	name = "Rahmat"
	fmt.Println(name)

	name = "Budi"
	fmt.Println(name)

	// without defined type variable
	var friendName = "Anto"
	fmt.Println(friendName)

	var age = 27
	fmt.Println(age)

	// without var
	country := "Indonesia"
	fmt.Println(country)

	country = "US"
	fmt.Println(country)

	// multiple variable
	var (
		myname    = "Rahmat"
		mycountry = "Indonesia"
	)

	fmt.Println(myname)
	fmt.Println(mycountry)
}
