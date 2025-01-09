package pokedex

import (
	"fmt"

	"VictorVolovik/go-pokedex/api"
)

type Pokedex struct {
	pokemons map[string]api.PokemonDetails
}

var UsersPokedex = NewPokedex()

func NewPokedex() *Pokedex {
	return &Pokedex{
		pokemons: make(map[string]api.PokemonDetails),
	}
}

func (pkdx *Pokedex) CheckPokemon(pokemonName string) (api.PokemonDetails, error) {
	pokemon, ok := pkdx.pokemons[pokemonName]
	if !ok {
		return api.PokemonDetails{}, fmt.Errorf("no pokemon found")
	}

	return pokemon, nil
}

func (pkdx *Pokedex) RecordPokemon(p api.PokemonDetails) {
	pkdx.pokemons[p.Name] = p
}
