package main

import (
	"fmt"

	"example.com/linq"
)

func main() {
	values := []int{1, 2, 3}
	c := linq.Channel(values).Where(func(value int) bool {
		return value%2 == 0
	})
	fmt.Println(<-c)
}
