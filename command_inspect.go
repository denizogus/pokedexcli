package main

import "fmt"

func commandInspect(cfg *config, param []string) error {
	if len(param) != 1 {
		return fmt.Errorf("wrong parameters")
	}

	pokemon, ok := cfg.Pokedex.Pokemons[param[0]]
	if !ok {
		return fmt.Errorf("You have not caught this pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}

	return nil
}
