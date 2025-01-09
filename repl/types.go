package repl

import (
	"VictorVolovik/go-pokedex/api"
	"VictorVolovik/go-pokedex/pokedex"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, params ...string) error
}

type config struct {
	nextQuery string
	prevQuery string
	apiClient api.Client
	pokedex   pokedex.Pokedex
}
