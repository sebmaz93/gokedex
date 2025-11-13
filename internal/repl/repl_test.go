package repl

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello     worlD  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "zookiDex Pokemon",
			expected: []string{"zookidex", "pokemon"},
		},
		{
			input:    "Pokeemoon",
			expected: []string{"pokeemoon"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		acLen := len(actual)
		exLen := len(c.expected)
		if acLen != exLen {
			t.Errorf("actual:%v, expected:%v, len missmatch.", acLen, exLen)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual:%v, expected:%v, value missmatch.", word, expectedWord)
			}
		}
	}
}
