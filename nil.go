package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Rahmat")
	fmt.Println(person)

	person2 := NewMap("")
	if person2 == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println(person2)
	}

}
