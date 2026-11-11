package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, param []string) error {
	if len(param) == 0 {
		return fmt.Errorf("You need a pokemon to catch...")
	}

	pokemon, err := cfg.client.GetPokemon(param[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	chance := 100 / float64(pokemon.BaseExperience+100)

	if rand.Float64() <= chance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.Pokedex.Pokemons[pokemon.Name] = pokemon
		return nil
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}
}
