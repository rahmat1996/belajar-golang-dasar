package main

import "fmt"

type Address1 struct {
	City, Province, Country string
}

func changeCountryToSingapore(address *Address1) {
	address.Country = "Singapore"
}

func main() {

	var address = Address1{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}

	changeCountryToSingapore(&address)

	fmt.Println(address)

}
