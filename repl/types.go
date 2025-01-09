package repl

import (
	"VictorVolovik/go-pokedex/api"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config, params ...string) error
}

type config struct {
	nextQuery string
	preQuery  string
	apiClient api.Client
}
