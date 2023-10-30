package main

import "fmt"

func sayHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	fmt.Println(sayHello(""))
	fmt.Println(sayHello("Rahmat"))
}
