package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type PokeClient struct {
	Client http.Client
}

func (p *PokeClient) GetLocationAreas(givenUrl *string) (*LocationAreas, error) {
	url := baseurl + "/location-area"
	if givenUrl != nil {
		url = *givenUrl
	}

	var locationareas *LocationAreas = &LocationAreas{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := p.Client.Do(request)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, locationareas); err != nil {
		return nil, err
	}

	return locationareas, nil
}
