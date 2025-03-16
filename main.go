package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/cmd/commands"
	"pokedexcli/internal"
	"pokedexcli/internal/pokecache"
	"pokedexcli/internal/pokeconfig"
	"time"
)

var registry map[string]pokeconfig.CliCommand

func main() {
	scanner := bufio.NewScanner(os.Stdout)

	registry = map[string]pokeconfig.CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commands.Exit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commands.Help,
		},
		"map": {
			Name:        "map",
			Description: "Get a list of 20 next pokemon locations",
			Callback:    commands.Map,
		},
		"mapb": {
			Name:        "map back",
			Description: "Get a list of 20 previous pokemon locations",
			Callback:    commands.MapBack,
		},
		"explore": {
			Name:        "explore",
			Description: "Get information about area",
			Callback:    commands.Explore,
		},
		"catch": {
			Name:        "catch",
			Description: "Attempt to catch a pokemon",
			Callback:    commands.Catch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a caught pokemon",
			Callback:    commands.Inspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "Show all pokemon you have caught",
			Callback:    commands.Pokedex,
		},
	}

	config := pokeconfig.Config{
		Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Previous: "",
		Cache:    pokecache.NewCache(nil, time.Duration(time.Duration.Minutes(10))),
		Registry: registry,
    CaughtPokemon: make(map[string]internal.Pokemon),
	}

	fmt.Println("Welcome to the Pokedex!")

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := internal.CleanInput(scanner.Text())


		if len(input) == 0 {
			commands.Exit(&config, nil)
		}

		inputCommand := input[0]

		command, ok := registry[inputCommand]

		if !ok {
			continue
		}

		err := command.Callback(&config, input)
		if err != nil {
			fmt.Print(err)
      commands.Exit(&config, nil)
		}
	}
}
