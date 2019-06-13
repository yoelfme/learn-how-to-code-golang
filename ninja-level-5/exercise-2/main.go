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

	persons := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for i, p := range persons {
		fmt.Println(i, p)
	}

}
