package main

import (
	"fmt"
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
			input:    "  hello  world  this is another test",
			expected: []string{"hello", "world", "this", "is", "another", "test"},
		},
		{
			input:    "charmander bulbasaur  ",
			expected: []string{"charmander", "bulbasaur"},
		},
	}
	for _, c := range cases {
		fmt.Printf("case: %v \n", c)
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf(`lengths don't match %v, vs %v`, actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			fmt.Printf("word: %v expectedWord: %v \n", word, expectedWord)
			if word != expectedWord {
				t.Errorf(`CleanInput %v, want: %v`, word, expectedWord)
			}
		}
	}
}
