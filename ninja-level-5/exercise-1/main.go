package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	favoriteFlavors []string
}

func main() {
	p1 := person{
		first: "Yoel",
		last: "Monzon",
		favoriteFlavors: []string{"Coco", "Pistacho"},
	}

	p2 := person{
		first: "Jhon",
		last: "Doe",
		favoriteFlavors: []string{"Vainilla", "Chocolate"},
	}

	for _, p := range []person{p1, p2} {
		fmt.Println(p.first, p.last)

		for _, fv := range p.favoriteFlavors {
			fmt.Printf("\t\t%s\n", fv)
		}
	}
}
