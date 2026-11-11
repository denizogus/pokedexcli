package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/denizogus/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("start with help")

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()

		cleanInput := cleanInput(input)
		if len(cleanInput) == 0 {
			continue
		}

		command := cleanInput[0]

		if c, ok := cfg.commands[command]; ok {
			if err := c.callback(cfg, cleanInput[1:]); err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Lists the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Lists the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explores the location. (> explore canalave-city-area)",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attemps to catch the pokemon. (> catch pikachu)",
			callback:    commandCatch,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	commands map[string]cliCommand
	client   pokeapi.PokeClient
	NextURL  *string
	PrevURL  *string
	Pokedex  Pokedex
}

type Pokedex struct {
	Pokemons map[string]pokeapi.Pokemon
}
