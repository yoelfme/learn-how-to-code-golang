package main

import (
	"fmt"
)

func main() {
	x := struct{
		first string
		age int
		friends map[string]int
	}{
		first: "Yoel",
		age: 23,
		friends: map[string]int{
			"Jhon": 50,
			"Pepe": 30,
		},
	}

	fmt.Println(x)
}
