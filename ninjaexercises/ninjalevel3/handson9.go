package main

import (
	"fmt"
)

func main() {
	favSport := "soccer"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "soccer":
		fmt.Println("kick a ball")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}
