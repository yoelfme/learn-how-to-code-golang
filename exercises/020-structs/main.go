package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person // this is an anonymous field, and we can have promoted fields in anonymous fields
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 32,
		},
		ltk: true,
	}

	p1 := person{
		first: "Miss",
		last: "Moneypenny",
		age: 27,
	}

	fmt.Println(sa1.person.first, sa1.first)
	fmt.Println(p1)
}
