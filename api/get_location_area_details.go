package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaDetails(areaIdOrName string) (LocationAreaDetails, error) {
	reqUrl := PokedexApiUrl + LocationAreaEndpoint + "/" + areaIdOrName

	var (
		data []byte
		err  error
	)

	if entry, exists := c.cache.Get(reqUrl); exists {
		data = entry
	} else {
		req, err := http.NewRequest("GET", reqUrl, nil)
		if err != nil {
			return LocationAreaDetails{}, fmt.Errorf("error creating request: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaDetails{}, fmt.Errorf("error making request: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode == 404 {
			return LocationAreaDetails{}, fmt.Errorf("location not found")
		}

		if res.StatusCode != 200 {
			return LocationAreaDetails{}, fmt.Errorf("not OK HTTP status: %s", res.Status)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaDetails{}, fmt.Errorf("error reading response body: %w", err)
		}

		c.cache.Add(reqUrl, data)
	}

	var locationAreaDetails LocationAreaDetails
	if err = json.Unmarshal(data, &locationAreaDetails); err != nil {
		return LocationAreaDetails{}, fmt.Errorf("error converting response data to json: %w", err)
	}

	return locationAreaDetails, nil
}
