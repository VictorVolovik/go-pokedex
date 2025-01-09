package repl

import (
	"fmt"
	"net/url"

	"VictorVolovik/go-pokedex/api"
)

// Get the next page of locations
func commandMapf(cfg *config, params ...string) error {
	locationAreas, err := cfg.apiClient.GetLocationAreas(cfg.nextQuery)
	if err != nil {
		return fmt.Errorf("error getting location areas, %w", err)
	}

	err = updateNextAndPrevUrls(&locationAreas, cfg)
	if err != nil {
		return err
	}

	for _, locationArea := range locationAreas.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}

	return nil
}

// Get the previous page of locations
func commandMapb(cfg *config, params ...string) error {
	locationAreas, err := cfg.apiClient.GetLocationAreas(cfg.prevQuery)
	if err != nil {
		return fmt.Errorf("error getting location areas, %w", err)
	}

	updateNextAndPrevUrls(&locationAreas, cfg)

	for _, locationArea := range locationAreas.Results {
		fmt.Printf("%s\n", locationArea.Name)
	}

	return nil
}

func updateNextAndPrevUrls(locationAreas *api.LocationAreas, cfg *config) error {
	next := locationAreas.Next

	if next != nil {
		parsedURL, err := url.Parse(*next)
		if err != nil {
			return fmt.Errorf("Error parsing next URL: %w", err)
		}
		cfg.nextQuery = parsedURL.RawQuery
	}

	previous := locationAreas.Previous
	if previous != nil {
		parsedURL, err := url.Parse(*previous)
		if err != nil {
			return fmt.Errorf("Error parsing previous URL: %w", err)
		}
		cfg.prevQuery = parsedURL.RawQuery
	}

	return nil
}
