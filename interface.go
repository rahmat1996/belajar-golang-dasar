package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello2(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person1 struct {
	Name string
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (person Person1) GetName() string {
	return person.Name
}

func main() {
	var rahmat Person1
	rahmat.Name = "Rahmat"
	sayHello2(rahmat)

	cat := Animal{
		Name: "Meong",
	}
	sayHello2(cat)

}
