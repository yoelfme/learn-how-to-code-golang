package main

import (
	"fmt"
)

// Print every rune code point of the uppercase alphabet three times
func main() {
	firstLetter := 65
	lastLetter := 90
	for times := 0; times < 3; times++ {
		fmt.Println(times + 1)
		for i := firstLetter; i <= lastLetter; i++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}
