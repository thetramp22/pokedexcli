package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := cleanInput((c.input))
		actualLen := len(actual)
		expectedLen := len(c.expected)
		if !reflect.DeepEqual(actualLen, expectedLen) {
			t.Errorf("expected slice length: %v, actual: %v", expectedLen, actualLen)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if !reflect.DeepEqual(word, expectedWord) {
				t.Errorf("expected word: %v, actual: %v", expectedWord, word)
			}
		}
	}
}
