package input

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "alrEEt?",
			expected: []string{"alreet?"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Bulbasaur iS kEWl",
			expected: []string{"bulbasaur", "is", "kewl"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("%v provides %d words, expected %v, with %d words", actual, len(actual), c.expected, len(c.expected))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("expected %s to equal %s", word, expectedWord)
			}
		}
	}
}
