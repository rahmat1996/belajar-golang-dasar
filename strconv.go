package main

import (
	"fmt"
	"strconv"
)

func main() {
	// change from string to boolean
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	// change from string to int
	number, err := strconv.ParseInt("1000000", 10, 32)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	// change from int to string
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("2000000")
	fmt.Println(valueInt)
}
