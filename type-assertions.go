package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {

	// using switch
	result2 := random()

	switch value := result2.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}

	result := random()
	fmt.Println(result) // interface{} type data
	resultString := result.(string)
	fmt.Println(resultString) // string type data
	resultInt := result.(int)
	fmt.Println(resultInt) // will throw panic because not type int

}
