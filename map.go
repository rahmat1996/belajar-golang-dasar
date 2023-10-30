package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Rahmat",
		"country": "Indonesia",
	}

	person["likes"] = "Cat"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["country"])

	book := map[string]string{
		"title":    "Learning Golang",
		"author":   "Anonim",
		"category": "programming",
	}

	fmt.Println(book)
	fmt.Println(len(book))

	//delete key
	delete(book, "category")
	fmt.Println(book)
	fmt.Println(len(book))
}
