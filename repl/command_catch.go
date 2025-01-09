package repl

import (
	"fmt"
	"math"
	"math/rand"

	"VictorVolovik/go-pokedex/api"
	"VictorVolovik/go-pokedex/pokedex"
)

func commandCatch(cfg *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Unspecified pokemon\n")
	}

	pokemonIdOrName := params[0]

	pokemonDetails, err := cfg.apiClient.GetPokemonDetails(pokemonIdOrName)
	if err != nil {
		return fmt.Errorf("error getting pokemon details: %w", err)
	}

	fmt.Printf("Throwing a Pokeball at %s\n", pokemonDetails.Name)

	wasCaught := catchPokemon(&pokemonDetails)

	if wasCaught {
		fmt.Printf("%s was caught!\n", pokemonDetails.Name)
		pokedex.UsersPokedex.RecordPokemon(pokemonDetails)
	} else {
		fmt.Printf("%s escaped!\n", pokemonDetails.Name)
	}

	return nil
}

// catchPokemon attempts to catch a pokemon based on its base experience (BaseExperience).
// The probability to catch a pokemon decreases as its BaseExperience increases.
//
// Formula:
//
//	difficultyModifier = int(math.Sqrt(float64(BaseExperience))) * 5
//	roll = rand.Intn(BaseExperience + difficultyModifier)
//	Success if roll >= BaseExperience
//
// Examples:
//
//	BaseExperience:  64  -> Roll: 64 + 8 * 5 = 104 (~38% chance to catch)
//	BaseExperience: 100  -> Roll: 100 + 10 * 5 = 150 (~33% chance to catch)
//	BaseExperience: 400  -> Roll: 400 + 20 * 5 = 500 (~20% chance to catch)
func catchPokemon(p *api.PokemonDetails) bool {
	difficultyModifier := int(math.Sqrt(float64(p.BaseExperience))) * 5
	roll := rand.Intn(p.BaseExperience + difficultyModifier)

	return roll >= p.BaseExperience
}
