package main

import (
	"fmt"

	"VictorVolovik/go-pokedex/repl"
)

func main() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' for assistance or 'exit' to quit.")

	err := repl.Repl()
	if err != nil {
		fmt.Println(err)
	}
}
