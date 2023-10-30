package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	rahmat := Man{Name: "Rahmat"}
	rahmat.Married()
	fmt.Println(rahmat.Name)
}
