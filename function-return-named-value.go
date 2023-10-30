package main

import "fmt"

func getData() (name string, age int) {
	name = "Rahmat"
	age = 99
	return
}

func main() {
	a, b := getData()
	fmt.Println(a)
	fmt.Println(b)
}
