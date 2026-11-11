package main

import (
	"fmt"
)

func commandExplore(cfg *config, param []string) error {
	if len(param) == 0 {
		return fmt.Errorf("you need a place to explore...")
	}

	fmt.Printf("Exploring %s...\n", param[0])

	location, err := cfg.client.GetLocationArea(param[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, encounters := range location.PokemonEncounters {
		fmt.Println(" - " + encounters.Pokemon.Name)
	}

	return nil
}
