package main

import "fmt"

func main() {
	const PI = 3.14

	// multiple constant declaration
	const (
		name    string = "Rahmat"
		country        = "Indonesia"
		age     int    = 27
	)

	fmt.Println(PI)

	fmt.Println(name)
	fmt.Println(country)
	fmt.Println(age)
}
