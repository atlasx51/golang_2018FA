package main

import (
	"fmt"
)

func main() {
	x := []int{1, 2, 44}
	for i, _ := range x {
		fmt.Println(i, "-", x[i])
	}
}
