package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
  scanner := bufio.NewScanner(os.Stdout)

  for {
    fmt.Print("Pokedex > ")
    scanner.Scan()
    input := scanner.Text()
    command := strings.ToLower(input)
    command = strings.Fields(command)[0]
    fmt.Printf("Your command was: %v\n", command)

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
