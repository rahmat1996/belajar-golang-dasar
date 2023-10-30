package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Rahmat")
	// helper.sayGoodMorning("Rahmat") // cannot access because this private function
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // cannot access too
}
