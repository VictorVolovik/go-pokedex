package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"VictorVolovik/go-pokedex/api"
	"VictorVolovik/go-pokedex/pokedex"
)

func Repl(apiClient api.Client) error {
	scnanner := bufio.NewScanner(os.Stdin)
	cfg := &config{
		apiClient: apiClient,
		pokedex:   *pokedex.UsersPokedex,
	}

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
			userParams := words[1:]
			handleCommand(userCommand, userParams, cfg)
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
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location id or name>",
			description: "List all Pokemon in specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon id or name>",
			description: "Attempt to catch specified Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Check caught Pokemon details",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "List all caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

// Try to execute user's command
func handleCommand(userCommand string, userParams []string, cfg *config) {
	command, exists := getCommands()[userCommand]
	if !exists {
		fmt.Println("Unknown command")
		return
	}

	err := command.callback(cfg, userParams...)
	if err != nil {
		fmt.Println(err)
	}
}
