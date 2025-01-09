package repl

import "fmt"

// Display a help information
func commandHelp(cfg *config, params ...string) error {
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	return nil
}
