package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := address1
	address2.City = "Bandung"

	fmt.Println(address1) // address1 city value not changed bacause data address2 is passing value not referennce
	fmt.Println(address2)

	address3 := Address{"Malang", "Jawa Timur", "Indonesia"}
	address4 := &address3
	address4.City = "Pacitan"

	fmt.Println(address3) // will changed because using pointer as pass reference on passing data
	fmt.Println(address4)

	address5 := Address{"Magelang", "Jawa Tengah", "Indonesia"}
	address6 := &address5
	address7 := &address5

	*address6 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}

	fmt.Println(address5)
	fmt.Println(address6)
	fmt.Println(address7)

	var address8 *Address = new(Address)
	fmt.Println(address8)

}
