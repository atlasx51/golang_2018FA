package main

import (
	"fmt"
)

func main() {

	name := "Sam"

	str := `html here ` + name + `more html`

	fmt.Println(str)

	// fmt.Println(os.Args[0])
	// fmt.Println(os.Args[1])
	// fmt.Println(os.Args[2])
	// fmt.Println(os.Args[3])
}
