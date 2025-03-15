package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Expected and actual slice differ in length, actual %v - expected %v", len(actual), len(c.expected))
			t.Fail()
		}

		for i := range actual {
			word := actual[i]
			expected := c.expected[i]

			if word != expected {
				t.Errorf("expeted word %s didn't match actual %s", expected, word)
				t.Fail()
			}
		}
	}
}
