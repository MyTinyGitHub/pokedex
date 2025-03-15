package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
  name string
  description string
  callback func() error
}

var registry map[string]cliCommand

func main() {
  scanner := bufio.NewScanner(os.Stdout)
  registry = map[string]cliCommand{
    "exit": {
      name: "exit", 
      description: "Exit the Pokedex",
      callback: exitCommand,
    },
    "help": {
      name: "help", 
      description: "Displays a help message",
      callback: helpCommand,
    },
  }


  fmt.Println("Welcome to the Pokedex!")

  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    input := scanner.Text()
    if input == "" {
      exitCommand()
    }
    input = strings.Fields(input)[0]
    input = strings.ToLower(input)
    command, ok := registry[input]

    if !ok {
      continue
    }

    command.callback()
  }
}

func helpCommand() error {
  for _, value := range registry {
      fmt.Printf("%v: %v\n", value.name, value.description)
  }

  return nil
}

func exitCommand() error {
  fmt.Println("Closing the Pokedex... Goodbye!")
  os.Exit(0)
  return nil
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
