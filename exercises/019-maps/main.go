package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James": 32,
		"Juana": 27,
	}

	fmt.Println(m)
	fmt.Println(m["James"])
	fmt.Println(m["Barnabas"])

	// Use to verify if a map has a key
	v, ok := m["Barnabas"]
	fmt.Println(v, ok)

	// this is call "comma ok" idiom
	if v, ok := m["James"]; ok {
		fmt.Println("James exists, it's age is:", v)
	}

	m["Yoel"] = 23

	for k, v := range m {
		fmt.Println(k, v)
	}

	// delete a key from a map
	delete(m, "Juana")

	fmt.Println(m)
}
