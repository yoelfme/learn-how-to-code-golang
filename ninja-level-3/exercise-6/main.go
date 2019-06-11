package main

import (
	"fmt"
)

func main() {
	x := "Yoel Monzon"

	if len(x)> 10 {
		fmt.Println("Your name is long")
	} else {
		fmt.Println("Your name is short")
	}
}
