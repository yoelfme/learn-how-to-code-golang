package main

import (
	"fmt"
)

func main() {
	x:= map[string][]string{
		"bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr": []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	x["yoel"] = []string{"Computer Science", "Videogames", "Food"}

	delete(x, "no_dr")

	for k, v := range x {
		fmt.Println(k, v)
	}

}
