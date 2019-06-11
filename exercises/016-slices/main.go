package main

import "fmt"

// a SLICE allows you to group together VALUES of the same TYPE

func main() {
	// x := type{values} -> composite literal
	x := []int{4, 5,6,7,8}
	fmt.Println(x, len(x), cap(x))

	// looping trough an slice
	for i, v := range x {
		fmt.Println(i, v)
	}

	// slicing a slice
	y := x[1:]
	z := x[1:3]
	fmt.Println(y, len(y), cap(y))
	fmt.Println(z, len(z), cap(z))

	// append a slice
}
