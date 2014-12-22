package main

import (
	"fmt"
)

func Greet(name string) {
	fmt.Println("Hello, " + name)
}

func GreetNames(names []string, suffix string) {
	for _, name := range names {
		Greet(name + suffix)
	}
}

func main() {
	names := []string{
		"Mark",
		"RJ",
		"Ryan",
		"Leo",
	}
	comm := make(chan string)

	go func() {
		GreetNames(names, " <C> ")
		comm <- "finished greeting names concurrently..."
	}()

	GreetNames(names, " <M> ")
	fmt.Println(<-comm)
}
