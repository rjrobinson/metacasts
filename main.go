package main

import (
	"fmt"
)

func Greet(name string) {
	fmt.Println("Hello, " + name)
}

func main() {
	names := []string{
		"Mark",
		"RJ",
		"Ryan",
		"Leo",
	}

	for _, name := range names {
		Greet(name)
	}
}
