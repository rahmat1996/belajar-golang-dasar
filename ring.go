package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	// iterate to assign data
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	fmt.Println(*data) // cannot use println to show real data

	//user iterate to show real data
	data.Do(func(value interface{}) {
		fmt.Println(value)
	})
}
