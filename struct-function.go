package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

func (person Person) sayGoodNight(name string) {
	fmt.Println("Good Night", name, "From", person.Name)
}

func (person Person) sayGoodAfternoon(name string) {
	fmt.Println("Good Afternoon", name, "From", person.Name)
}

func main() {

	rahmat := Person{
		Name:    "Rahmat",
		Address: "Indonesia",
		Age:     99,
	}

	rahmat.sayGoodNight("Budi")
	rahmat.sayGoodAfternoon("Anto")

}
