package main

import "fmt"

func main() {

	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Not June"
	fmt.Println(slice1)

	slice1[0] = "Not May"
	fmt.Println(months)

	var days = [...]string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}

	fmt.Println(days)

	var daySlice1 = days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "New Saturday"
	daySlice1[1] = "New Sunday"
	fmt.Println(daySlice1)
	fmt.Println(days)

	var daySlice2 = append(daySlice1, "New Day")

	fmt.Println(daySlice2) // will return new array

	daySlice2[0] = "Old day"
	fmt.Println(daySlice2)
	fmt.Println(daySlice1)
	fmt.Println(days)

	//make slice
	newSlice := make([]string, 2, 5)

	newSlice[0] = "Rahmat"
	newSlice[1] = "Budi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)

	fmt.Println(copySlice)

	arrayExample := [...]int{1, 2, 3, 4, 5}
	sliceExample := []int{1, 2, 3, 4, 5}

	fmt.Println(arrayExample)
	fmt.Println(sliceExample)

}
