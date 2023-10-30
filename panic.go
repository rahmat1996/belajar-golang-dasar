package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi selesai")
}

func runApp(isError bool) {
	defer endApp()
	if isError {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(false)
	runApp(true)
}
