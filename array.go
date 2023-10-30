package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Rahmat"
	names[1] = "Budi"
	names[2] = "Anto"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		80,
		70,
		90,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var statuses [10]string

	fmt.Println(len(statuses))

}
