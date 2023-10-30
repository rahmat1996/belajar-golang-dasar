package main

import "fmt"

type Filter func(string) string // declaration type for filter

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Budi", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
