package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config, param []string) error {
	if len(param) != 0 {
		fmt.Errorf("wrong parameters")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
