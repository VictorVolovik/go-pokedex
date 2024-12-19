package repl

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "test string ",
			expected: []string{"test", "string"},
		},
		{
			input:    "Title Cased",
			expected: []string{"title", "cased"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		actualLen := len(actual)
		expectedLen := len(c.expected)
		if actualLen != expectedLen {
			t.Errorf("lengths are not equal")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("words at index %d are not equal", i)
			}
		}
	}
}
