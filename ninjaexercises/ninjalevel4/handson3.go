package main

import (
	"fmt"
)

func main() {
	x := []int{12, 13, 14, 15, 16, 17, 18, 19, 20, 21}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)
}
