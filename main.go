package main

import (
	"fmt"

	"VictorVolovik/go-pokedex/repl"
)

func main() {
	fmt.Println("Welcome to the Pokedex!")

	err := repl.Repl()
	if err != nil {
		fmt.Println(err)
	}
}
