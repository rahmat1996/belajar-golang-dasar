package main

import "fmt"

func logging() {
	fmt.Println("Selesai memangil function")
}

func runApplication(value int) {
	fmt.Println("Run Application")
	result := 10 / value
	fmt.Println("Result", result)
	logging()
}

func runApplicationDefer(value int) {
	defer logging()
	fmt.Println("Run Application Defer")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(10)
	runApplicationDefer(0)
}
