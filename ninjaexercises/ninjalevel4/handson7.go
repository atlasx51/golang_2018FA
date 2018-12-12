package main

import (
	"fmt"
)

func main() {
	xs1 := []string{"Chuky", "Lozano", "PSV"}
	xs2 := []string{"PSG", "LYON", "Champions League"}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)

	for i, xs := range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
