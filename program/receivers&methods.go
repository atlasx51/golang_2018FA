package main

import (
	"fmt"
)

type person struct {
	last string
	age  int
}

func main() {
	p1 := person{"Bubulubu", 24}
	fmt.Println(p1)

	p1.boo()

}

func (p person) boo() {
	fmt.Println(p.last, "is walking.")
}

// func (receiver) identifier(parameters) (returns) {}
