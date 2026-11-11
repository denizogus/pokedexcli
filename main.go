package main

import (
	"net/http"

	"github.com/denizogus/pokedexcli/internal/pokeapi"
)

func main() {
	pokeclient := pokeapi.PokeClient{
		Client: http.Client{},
	}

	cfg := &config{
		commands: getCommands(),
		client:   pokeclient,
	}

	startRepl(cfg)
}
