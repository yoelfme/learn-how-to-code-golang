package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
