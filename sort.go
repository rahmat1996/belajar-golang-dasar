package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

// function interface from sort
func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i int, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i int, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {

	var users []User = []User{
		{Name: "Rahmat", Age: 99},
		{Name: "Budi", Age: 95},
		{Name: "Anto", Age: 90},
		{Name: "Eko", Age: 89},
	}

	fmt.Println(users)
	sort.Sort(UserSlice(users))
	fmt.Println(users)

}
