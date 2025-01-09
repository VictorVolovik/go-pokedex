package repl

import "fmt"

func commandExplore(cfg *config, params ...string) error {
	if len(params) == 0 {
		return fmt.Errorf("Unspecified location area\n")
	}

	areaIdOrName := params[0]

	locationAreaDetails, err := cfg.apiClient.GetLocationAreaDetails(areaIdOrName)
	if err != nil {
		return fmt.Errorf("error getting location area details, %w", err)
	}

	fmt.Printf("Exploring %s\n", locationAreaDetails.Name)

	for _, pokemonEncounter := range locationAreaDetails.PokemonEncounters {
		fmt.Printf("  - %s\n", pokemonEncounter.Pokemon.Name)
	}

	return nil
}
