package repl

import (
	"fmt"
	"os"
)

// Exit Pokedex
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
