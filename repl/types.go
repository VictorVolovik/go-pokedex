package repl

type cliCommand struct {
	name        string
	description string
	callback    func() error
}
