package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()

	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2023, 10, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02" // LOL, why not 'YYYY-MM-DD' ?
	parse, _ := time.Parse(layout, "2023-10-30")
	fmt.Println(parse)

}
