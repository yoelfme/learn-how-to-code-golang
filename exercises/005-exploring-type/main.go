package main

import "fmt"

var y = 42

// DECLARE the VARIABLE with the IDENTIFIER "z"
// is of TYPE string and I ASSIGN the VALUE "shaken , not stirred"
var z = "Shaken, not stirred"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)

}
