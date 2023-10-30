package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.44444444444444444444444444444449))
	fmt.Println(math.Round(1.4999999999999999999))
	fmt.Println(math.Round(1.5))

	fmt.Println(math.Floor(1.99999999999999999999999999999999999999))

	fmt.Println(math.Ceil(1.000000000000000000000000000000000000009))

	fmt.Println(math.Max(10, 20))
	fmt.Println(math.Min(10, 20))
}
