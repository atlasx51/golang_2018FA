package main

import (
	"fmt"
)

func main() {
	x := map[string]int{"Bob": 45, "Thierry": 25, "Antoine": 27}
	for k, v := range x {
		fmt.Println(k, v)
	}
}
