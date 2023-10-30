package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var rahmat Customer
	rahmat.Name = "Rahmat"
	rahmat.Address = "Indonesia"
	rahmat.Age = 99

	fmt.Println(rahmat)
	fmt.Println(rahmat.Name)
	fmt.Println(rahmat.Address)
	fmt.Println(rahmat.Age)

	budi := Customer{
		Name:    "Budi",
		Address: "USA",
		Age:     91,
	}
	fmt.Println(budi)

	// not recomended
	anto := Customer{"Anto", "Singapore", 90}
	fmt.Println(anto)
}
