package main

import "fmt"

func main() {

	var name1 = "Rahmat"
	var name2 = "Rahmat"

	var result = name1 == name2
	fmt.Println(result)

	var name3 = "Budi"
	result = name1 == name3
	fmt.Println(result)

	var value1 = 100
	var value2 = 50

	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 >= value2)
	fmt.Println(value1 <= value2)
}
