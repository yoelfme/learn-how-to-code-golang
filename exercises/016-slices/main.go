package main

import "fmt"

// a SLICE allows you to group together VALUES of the same TYPE

func main() {
	// x := type{values} -> composite literal
	x := []int{4, 5,6,7,8}
	fmt.Println("original x", x, "length:", len(x), "capacity:", cap(x))

	// looping trough an slice
	for i, v := range x {
		fmt.Println(i, v)
	}

	// slicing a slice
	y := x[1:]
	z := x[1:3]
	fmt.Println("x sliced from item 1 to the latest item:", y,  "length:", len(y), "capacity:", cap(y))
	fmt.Println("x sliced from item 1 to item 2", z, "length:", len(z), "capacity:", cap(z))

	// append a slice
	x = append(x, 77, 88, 99)
	fmt.Println("x appended:", x)

	a := []int{100, 101, 102,103}
	x = append(x, a...)
	fmt.Println("a:", a)
	fmt.Println("a appended to x:", x)

	// delete from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println("x without the old position 2 and 3:", x)
}
