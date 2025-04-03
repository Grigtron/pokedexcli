package main

import (
	"testing"
	
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input      string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "gEoDuDe  MachAmp   pikaCHU   ",
			expected: []string{"geodude", "machamp", "pikachu"},
		},
		{
			input: "     diglet dugtriO   ",
			expected: []string{"diglet", "dugtrio"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Len of %d not equal to %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual [i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Word at index %d: %s not equal to Expected Word: %s", i, word, expectedWord)
			}
		}
	}
}