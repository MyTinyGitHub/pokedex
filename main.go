package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/commands"
	"pokedexcli/internal/pokecache"
	"pokedexcli/internal/pokeconfig"
	"strings"
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
		input := scanner.Text()

		if input == "" {
			commands.ExitCommand(&config)
		}

		input = strings.Fields(input)[0]
		input = strings.ToLower(input)
		command, ok := registry[input]

		if !ok {
			continue
		}

		err := command.Callback(&config)
		if err != nil {
			fmt.Print(err)
			commands.ExitCommand(&config)
		}
	}
}

func cleanInput(text string) []string {
	trimmed := strings.Trim(text, " ")
	split := strings.Split(trimmed, " ")
	result := make([]string, 0)

	for _, word := range split {
		result = append(result, strings.ToLower(word))
	}

	return result
}
