package main

import "fmt"

func main() {
	type NoKTP string
	type Contract bool

	var myNoKTP NoKTP = "1234567890"
	var myStatus Contract = true

	fmt.Println(myNoKTP)
	fmt.Println(myStatus)
}
