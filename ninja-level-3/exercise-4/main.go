package main

import (
	"fmt"
)

// Have it print out the years you have been alive.
func main() {
	bd := 1995
	for {
		fmt.Println(bd)
		bd++

		if bd > 2019 {
			break
		}
	}
}
