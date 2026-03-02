package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{input: " hello world ", expected: []string{"hello", "world"}},
		{input: "   foo bar baz   ", expected: []string{"foo", "bar", "baz"}},
		{input: "   singleword   ", expected: []string{"singleword"}},
		{input: "   multiple   spaces   between   words   ", expected: []string{"multiple", "spaces", "between", "words"}},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Expected %d words but got %d", len(c.expected), len(actual))
			return
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected word '%s' but got '%s'", expectedWord, word)
			}
		}
	}
}
