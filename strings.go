package main

import (
	"fmt"
	"strings"
)

func main() {
	//contains
	fmt.Println(strings.Contains("Rahmat Belajar Golang", "Rahmat"))
	fmt.Println(strings.Contains("Rahmat Belajar Golang", "PHP"))

	//split
	fmt.Println(strings.Split("Rahmat Belajar Golang", " "))

	//to lower case
	fmt.Println(strings.ToLower("Rahmat Belajar Golang"))

	//to upper case
	fmt.Println(strings.ToUpper("Rahmat Belajar Golang"))

	// trim
	fmt.Println(strings.Trim("     Rahmat Belajar Golang     ", " "))

	// replace all
	fmt.Println(strings.ReplaceAll("Rahmat Belajar PHP", "PHP", "Golang"))
}
