package main

import "fmt"

func commandPokedex(cfg *config, param []string) error {
	if len(param) != 0 {
		return fmt.Errorf("wrong parameters")
	}

	fmt.Println("Your Pokedex: ")
	for _, val := range cfg.Pokedex.Pokemons {
		fmt.Printf(" - %s\n", val.Name)
	}

	return nil
}
