package internal

import (
	"strings"
)

func CleanInput(text string) []string {
	result := strings.Fields(text)
	return result
}
