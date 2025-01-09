package repl

import "fmt"

// Display a list of caught Pokemon from Pokedex
func commandPokedex(cfg *config, params ...string) error {
	fmt.Println("Your Pokedex:")

	for _, pokemonName := range cfg.pokedex.GetPokemonNames() {
		fmt.Printf("-- %s\n", pokemonName)
	}

	return nil
}
