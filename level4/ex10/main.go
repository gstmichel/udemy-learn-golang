package main

import (
	"fmt"
)

func main() {
	x := map[string][]string{
		`bond_james`:{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`:{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:{`Being evil`, `Ice cream`, `Sunsets`},
	}

	delete(x, `no_dr`)

	for k, v := range x {
		fmt.Println(k)
		for idx, s := range v {
			fmt.Printf("\t%v\t%v\n", idx, s)
		}
	}
}
