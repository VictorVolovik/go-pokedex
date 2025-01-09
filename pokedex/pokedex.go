package pokedex

import (
	"fmt"

	"VictorVolovik/go-pokedex/api"
)

type Pokedex struct {
	records map[string]api.PokemonDetails
}

var UsersPokedex = NewPokedex()

func NewPokedex() *Pokedex {
	return &Pokedex{
		records: make(map[string]api.PokemonDetails),
	}
}

func (pkdx *Pokedex) CheckPokemon(pokemonName string) (api.PokemonDetails, error) {
	pokemon, ok := pkdx.records[pokemonName]
	if !ok {
		return api.PokemonDetails{}, fmt.Errorf("no pokemon found")
	}

	return pokemon, nil
}

func (pkdx *Pokedex) RecordPokemon(p api.PokemonDetails) {
	pkdx.records[p.Name] = p
}

func (pkdx *Pokedex) GetPokemonNames() []string {
	names := make([]string, len(pkdx.records))

	i := 0
	for _, pokemonDetails := range pkdx.records {
		names[i] = pokemonDetails.Name
		i++
	}

	return names
}
