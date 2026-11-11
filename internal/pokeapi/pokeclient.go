package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/denizogus/pokedexcli/internal/pokecache"
)

type PokeClient struct {
	Client http.Client
	Cache  *pokecache.Cache
}

func (p *PokeClient) GetLocationAreas20(givenUrl *string) (*LocationAreas20, error) {
	url := baseurl + "/location-area"
	if givenUrl != nil {
		url = *givenUrl
	}

	var locationareas *LocationAreas20 = &LocationAreas20{}

	if data, ok := p.Cache.Get(url); ok {
		if err := json.Unmarshal(data, locationareas); err != nil {
			return nil, err
		}
		return locationareas, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := p.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, locationareas); err != nil {
		return nil, err
	}

	p.Cache.Add(url, data)

	return locationareas, nil
}

func (p *PokeClient) GetLocationArea(name string) (*LocationArea, error) {
	url := baseurl + "/location-area/" + name
	var locationarea *LocationArea = &LocationArea{}

	if data, ok := p.Cache.Get(url); ok {
		if err := json.Unmarshal(data, locationarea); err != nil {
			return nil, err
		}
		return locationarea, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	response, err := p.Client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return nil, fmt.Errorf("Unknown location. (non-ok response code)")
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, locationarea); err != nil {
		return nil, err
	}

	p.Cache.Add(url, data)
	return locationarea, nil
}
