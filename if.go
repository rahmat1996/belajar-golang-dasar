package main

import "fmt"

func main() {

	var myName string = "Rahmat"

	// condition true
	if myName == "Rahmat" {
		fmt.Println("Hello Rahmat")
	}

	var friendName string = "Anto"

	// condition false
	if friendName == "Budi" {
		fmt.Println("My Friend Name Budi")
	}

	var exam = 50

	// use else
	if exam >= 70 {
		fmt.Println("Congrats!")
	} else {
		fmt.Println("Please try again.")
	}

	var season = "Winter"

	// use else if
	if season == "Spring" {
		fmt.Println("Season is Spring")
	} else if season == "Summer" {
		fmt.Println("Season is Summer")
	} else if season == "Fall" {
		fmt.Println("Season is Fall")
	} else if season == "Winter" {
		fmt.Println("Season is Winter")
	} else {
		fmt.Println("Season unknown")
	}

	//short expression
	var name = "Rahmat"

	if length := len(name); length > 5 {
		fmt.Println("Name too long")
	} else {
		fmt.Println("Okay, its good")
	}

	name = "Budi"

	if length := len(name); length > 5 {
		fmt.Println("Name too long")
	} else {
		fmt.Println("Okay, its good")
	}

}
