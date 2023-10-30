package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

func main() {
	sayHello()
	sayGreeting()
	sayHello()
}

func sayGreeting() {
	fmt.Println("Good morning")
}
