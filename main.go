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
			Callback:    commands.ExitCommand,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commands.HelpCommand,
		},
		"map": {
			Name:        "map",
			Description: "Get a list of 20 next pokemon locations",
			Callback:    commands.MapCommand,
		},
		"mapb": {
			Name:        "map back",
			Description: "Get a list of 20 previous pokemon locations",
			Callback:    commands.MapBackCommand,
		},
		"explore": {
			Name:        "explore",
			Description: "Get information about area",
			Callback:    commands.ExploreCommand,
		},
	}

	config := pokeconfig.Config{
		Next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		Previous: "",
		Cache:    pokecache.NewCache(nil, time.Duration(time.Duration.Minutes(10))),
		Registry: registry,
	}

	fmt.Println("Welcome to the Pokedex!")

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := internal.CleanInput(scanner.Text())


		if len(input) == 0 {
      stringlings := [2]string{"a", "pastoria-city-area"}
      commands.ExploreCommand(&config,stringlings[:])
			commands.ExitCommand(&config, nil)
		}

		inputCommand := input[0]

		command, ok := registry[inputCommand]

		if !ok {
			continue
		}

		err := command.Callback(&config, input)
		if err != nil {
			fmt.Print(err)
      commands.ExitCommand(&config, nil)
		}
	}
}
