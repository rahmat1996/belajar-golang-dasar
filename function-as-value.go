package main

import "fmt"

func getGoodMorning(name string) string {
	return "Good Morning " + name
}

func main() {
	sayGoodMorning := getGoodMorning
	fmt.Println(sayGoodMorning("Rahmat"))
	fmt.Println(getGoodMorning("Rahmat"))
}
