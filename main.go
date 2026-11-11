package main

import (
	"net/http"
	"time"

	"github.com/denizogus/pokedexcli/internal/pokeapi"
	"github.com/denizogus/pokedexcli/internal/pokecache"
)

func main() {
	pokeclient := pokeapi.PokeClient{
		Client: http.Client{},
		Cache:  pokecache.NewCache(5 * time.Minute),
	}

	cfg := &config{
		commands: getCommands(),
		client:   pokeclient,
		Pokedex: Pokedex{
			Pokemons: map[string]pokeapi.Pokemon{},
		},
	}

	startRepl(cfg)
}
