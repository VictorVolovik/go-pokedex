package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Repl() error {
	scnanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		hasText := scnanner.Scan()
		if hasText {
			input := scnanner.Text()
			words := cleanInput(input)
			if len(words) == 0 {
				continue
			}
			command := words[0]
			fmt.Printf("Your command was: %s\n", command)
		} else {
			err := scnanner.Err()
			if err != nil {
				return fmt.Errorf("failed to scan user input: %w", err)
			}
		}
	}
}

// Split, lowercase and trim user's input
func cleanInput(text string) []string {
	lowerCased := strings.ToLower(text)
	words := strings.Fields(lowerCased)
	return words
}
