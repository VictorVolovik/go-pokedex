package api

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreas struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationAreaDetails struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon Pokemon `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
