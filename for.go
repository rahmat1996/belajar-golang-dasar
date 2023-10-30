package main

import "fmt"

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	//short version
	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan (short) ke", i)
	}

	names := []string{
		"Rahmat",
		"Budi",
		"Anto",
		"Joko",
		"Eko",
	}

	//for range
	for i, value := range names {
		fmt.Println("Index", i, "=", value)
	}

	//for range without index (use underscore to bypass variable index)
	for _, value := range names {
		fmt.Println(value)
	}

	//for range on map
	var person = map[string]string{
		"name":    "Rahmat",
		"country": "Indonesia",
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}

}
