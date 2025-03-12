package main

import (
	"fmt"
	"strings"
)

func main() {
    fmt.Print("Hello, World!")
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
