package pokedex

import (
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

func (pkdx *Pokedex) CheckPokemon(pokemonName string) api.PokemonDetails {
	pokemon, ok := pkdx.pokemons[pokemonName]
	if !ok {
		return api.PokemonDetails{}
	}

	return pokemon
}

func (pkdx *Pokedex) RecordPokemon(p api.PokemonDetails) {
	pkdx.pokemons[p.Name] = p
}
