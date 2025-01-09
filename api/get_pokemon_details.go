package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(pokemonIdOrName string) (PokemonDetails, error) {
	reqUrl := PokedexApiUrl + PokemonEndpoint + "/" + pokemonIdOrName

	var (
		data []byte
		err  error
	)

	if entry, exists := c.cache.Get(reqUrl); exists {
		data = entry
	} else {
		req, err := http.NewRequest("GET", reqUrl, nil)
		if err != nil {
			return PokemonDetails{}, fmt.Errorf("error creating request: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return PokemonDetails{}, fmt.Errorf("error making request: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode == 404 {
			return PokemonDetails{}, fmt.Errorf("pokemon not found")
		}

		if res.StatusCode != 200 {
			return PokemonDetails{}, fmt.Errorf("not OK HTTP status: %s", res.Status)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return PokemonDetails{}, fmt.Errorf("error reading response body: %w", err)
		}

		c.cache.Add(reqUrl, data)
	}

	var pokemonDetails PokemonDetails
	if err = json.Unmarshal(data, &pokemonDetails); err != nil {
		return PokemonDetails{}, fmt.Errorf("error converting response data to json: %w", err)
	}

	return pokemonDetails, nil
}
