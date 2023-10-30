package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runAppWithRecover(isError bool) {
	defer endApp()
	if isError {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runAppWithRecover(false)
	runAppWithRecover(true)
	fmt.Println("Rahmat")
}
