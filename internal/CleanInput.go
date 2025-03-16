package internal

import (
  "strings"
)

func CleanInput(text string) []string {
	result := strings.Fields(text)
	//split := strings.Split(trimmed, " ")
	//result := make([]string, 0)

	//for _, word := range split {
		//result = append(result, strings.ToLower(word))
	//}

	return result
}
