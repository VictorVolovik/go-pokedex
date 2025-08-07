package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const DefaultQueryUrl = "offset=0&limit=20"

func (c *Client) GetLocationAreas(queryUrl string) (LocationAreas, error) {
	reqUrl := PokedexApiUrl + LocationAreaEndpoint
	if len(queryUrl) > 0 {
		reqUrl += "?" + queryUrl
	} else {
		reqUrl += "?" + DefaultQueryUrl
	}

	var (
		data []byte
		err  error
	)

	if entry, exists := c.cache.Get(reqUrl); exists {
		data = entry
	} else {
		req, err := http.NewRequest("GET", reqUrl, nil)
		if err != nil {
			return LocationAreas{}, fmt.Errorf("error creating request: %w", err)
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreas{}, fmt.Errorf("error making request: %w", err)
		}
		defer res.Body.Close()

		if res.StatusCode != 200 {
			return LocationAreas{}, fmt.Errorf("not OK HTTP status: %s", res.Status)
		}

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreas{}, fmt.Errorf("error reading response body: %w", err)
		}

		c.cache.Add(reqUrl, data)
	}

	var locationAreas LocationAreas
	if err = json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreas{}, fmt.Errorf("error converting response data to json: %w", err)
	}
	return locationAreas, nil
}
