package main

import "fmt"

type yalo int

var x yalo

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
