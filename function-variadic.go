package main

import "fmt"

func sumAll(numbers ...int) int {
	var total int = 0
	for _, value := range numbers {
		total += value
	}
	return total
}

func main() {
	total := sumAll(50, 50, 30, 20, 45, 40, 80, 2, 90, 90)
	fmt.Println(total)

	slice := []int{
		90, 90, 70, 70, 50, 30,
	}

	fmt.Println(sumAll(slice...))
}
