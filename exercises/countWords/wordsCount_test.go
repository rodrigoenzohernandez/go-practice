package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordCount(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{
			input: "My name is Rodrigo and I am learning Go!",
			expected: map[string]int{
				"My":       1,
				"name":     1,
				"is":       1,
				"Rodrigo":  1,
				"and":      1,
				"I":        1,
				"am":       1,
				"learning": 1,
				"Go!":      1,
			},
		},
		{
			input: "Go Js Js Java Java Java C# C# C#",
			expected: map[string]int{
				"Go":   1,
				"Js":   2,
				"Java": 3,
				"C#":   3,
			},
		},
		{
			input:    "",
			expected: map[string]int{},
		},
	}

	for _, test := range tests {
		result := WordsCount(test.input)
		assert.Equal(t, test.expected, result, "WordCount(%q) should match expected value", test.input)
	}
}
