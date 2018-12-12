package main

import (
	"fmt"
)

type person struct {
	fname   string
	age     int
	favFood []string
}

func main() {
	p1 := person{
		"Bond",
		32,
		[]string{"tacos", "pizza", "burgers", "orange chicken"},
	}
	fmt.Println(p1.favFood)
}
