package main

import (
	"fmt"
	"time"

	"VictorVolovik/go-pokedex/api"
	"VictorVolovik/go-pokedex/repl"
)

func main() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Type 'help' for assistance or 'exit' to quit.")

	apiClient := api.NewClient(5*time.Second, time.Minute*5)

	err := repl.Repl(apiClient)
	if err != nil {
		fmt.Println(err)
	}
}
