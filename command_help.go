package main

import "fmt"

func commandHelp(cfg *config, param []string) error {
	if len(param) != 0 {
		fmt.Errorf("wrong parameters")
	}

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, c := range cfg.commands {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}
