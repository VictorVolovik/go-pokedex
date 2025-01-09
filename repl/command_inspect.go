package repl

import (
	"VictorVolovik/go-pokedex/api"
	"fmt"
)

// Show caught Pokemon details
func commandInspect(cfg *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Unspecified pokemon\n")
	}

	pokemonName := params[0]

	pokemonDetails, err := cfg.pokedex.CheckPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("you have not caught that pokemon")
	}

	printPokemonDetails(&pokemonDetails)

	return nil
}

func printPokemonDetails(p *api.PokemonDetails) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)

	fmt.Printf("Stats:\n")
	for _, stat := range p.Stats {
		fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Printf("Types:\n")
	for _, pType := range p.Types {
		fmt.Printf("  - %s\n", pType.Type.Name)
	}
}
