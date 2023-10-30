package helper

import "fmt"

var version = 1
var Application = "Belajar Golang Dasar"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

/**
This function cannot access because private function, first character function name is lowerCase.
For public access first character function name must UpperCase
*/
func sayGoodMorning(name string) {
	fmt.Println("Good Morning", name)
}
