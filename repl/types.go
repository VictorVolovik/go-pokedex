package repl

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *config) error
}

type config struct {
	Next     string
	Previous string
}
