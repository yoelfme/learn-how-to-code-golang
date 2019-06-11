package main

import (
	"fmt"
	"runtime"
)

var x byte
var y uint8

func main() {

	x = 2
	y = 3

	floatingNumber := 4.85

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%.2f", floatingNumber)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
