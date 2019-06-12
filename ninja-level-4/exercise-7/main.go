package main

import (
	"fmt"
)

func main() {
	x := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloooooo, James."},
	}

	for _, y := range x {
		for _, z := range y {
			fmt.Println(z)
		}
	}
}
