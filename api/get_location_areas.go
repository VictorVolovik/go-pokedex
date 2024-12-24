package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationAreas(queryUrl string) (LocationAreas, error) {
	reqUrl := PokedexApiUrl + LocationAreaEndpoint
	if len(queryUrl) > 0 {
		reqUrl += "?" + queryUrl
	}

	res, err := http.Get(reqUrl)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error  creating request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreas{}, fmt.Errorf("error reading response body: %w", err)
	}

	var locationAreas LocationAreas
	if err = json.Unmarshal(data, &locationAreas); err != nil {
		return LocationAreas{}, fmt.Errorf("error converting response data to json: %w", err)
	}
	return locationAreas, nil
}
