package main

import "fmt"

func sayHelloTo(name string) {
	fmt.Println("Hello", name)
}

func main() {
	sayHelloTo("Rahmat")
	sayHelloToFriend("Budi", "Anto")
}

func sayHelloToFriend(firstFriend string, secondFriend string) {
	fmt.Println("Hello", firstFriend, "And", secondFriend)
}
