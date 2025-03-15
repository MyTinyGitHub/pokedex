package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"pokedexcli/internal/pokecache"
	"strings"
	"time"
)

type cliCommand struct {
  name string
  description string
  callback func(*Config) error
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
    "map": {
      name: "map",
      description: "Get a list of 20 next pokemon locations",
      callback: mapCommand,
    },
    "mapb": {
      name: "map back",
      description: "Get a list of 20 previous pokemon locations",
      callback: mapBackCommand,
    },
  }

  config := Config{
    Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
    Previous: "",
    Cache: pokecache.NewCache(nil, time.Duration(time.Duration.Minutes(10))),
  }


  fmt.Println("Welcome to the Pokedex!")

  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    input := scanner.Text()

    if input == "" {
      exitCommand(&config)
    }

    input = strings.Fields(input)[0]
    input = strings.ToLower(input)
    command, ok := registry[input]

    if !ok {
      continue
    }

    err := command.callback(&config)
    if err != nil {
      fmt.Print(err)
      exitCommand(&config)
    }
  }
}

type Config struct {
  Next string
  Previous string
  Cache pokecache.Cache
}

type LocationArea struct {
  Count int `json:"count"`
  Next string `json:"next"`
  Previous string `json:"previous"`
  Location []LocationResults `json:"results"`
}

type LocationResults struct {
  Name string `json:"name"`
  Url string `json:"url"`
}


func helpCommand(config *Config) error {
  for _, value := range registry {
      fmt.Printf("%v: %v\n", value.name, value.description)
  }

  return nil
}

func getDataFromUrl(url string) ([]byte, error) {
  res, err := http.Get(url)
  if err != nil {
    return nil, err
  }

  defer res.Body.Close()

  return io.ReadAll(res.Body)
}

func getData(url string, c *Config) ([]byte, error) {
  value, ok := c.Cache.Get(url)
  if !ok {
    data, err := getDataFromUrl(url)
    if err != nil {
      return nil, err
    }

    c.Cache.Add(url, data)

    return data, nil
  }
  fmt.Println("got from cache")
  return value, nil
}


func getLocationAreas(url string, c *Config) (LocationArea, error) {
  value, err := getData(url, c)
  if err != nil {
    return LocationArea{}, err
  }

  var locations LocationArea

  err = json.Unmarshal(value, &locations)
  if err != nil {
    return LocationArea{}, fmt.Errorf("unable to unmarshal: %v", err)
  }

  return locations, nil
}

func mapCommand(config *Config) error {
  locations, err := getLocationAreas(config.Next, config)
  if err != nil {
    return err
  }

  config.Next = locations.Next
  config.Previous = locations.Previous

  for _, location := range locations.Location {
    fmt.Println(location.Name)
  }

  return nil
}

func mapBackCommand(config *Config) error {
  if(config.Previous == "") {
    fmt.Println("You're on the first page")
    return nil
  }

  locations, err := getLocationAreas(config.Previous, config)
  if err != nil {
    return err
  }

  config.Next = locations.Next
  config.Previous = locations.Previous

  for _, location := range locations.Location {
    fmt.Println(location.Name)
  }

  return nil
}

func exitCommand(config *Config) error {
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
