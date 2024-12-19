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
			userCommand := words[0]
			handleCommand(userCommand)
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

// Get a map of commands
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

// Try to execute user's command
func handleCommand(userCommand string) {
	command, exists := getCommands()[userCommand]
	if !exists {
		fmt.Println("Unknown command")
		return
	}

	err := command.callback()
	if err != nil {
		fmt.Println(err)
	}
}
