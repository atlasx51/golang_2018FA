package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`lionel_messi`: []string{`goals`, `goldenboot`, `bcn`},
		`rafa_marquez`: []string{`bcn`, `gdl`, `kaiser`},
		`dr_d`:         []string{`good`, `tacos`, `al pastor`},
	}

	// fmt.Println(m)

	m[`pavel_pardo`] = []string{`stuttgart`, `atlas`, `ame`}

	delete(m, `dr_d`)

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
