package main

import (
	"fmt"
)

func main() {
	x := "Lionel Messi"

	if x == "Lionel Messi" {
		fmt.Println(x)
	} else if x == "Cristiano Ronaldo" {
		fmt.Println("Juventus", x)
	} else {
		fmt.Println("neither")
	}
}
