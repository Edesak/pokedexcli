package main

import (
	"testing"
)

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"

	SuccessIcon = "✔"
	FailureIcon = "✘"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Trailing and Leading Spaces",
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			name:     "Empty String",
			input:    "",
			expected: []string{},
		},
		{
			name:     "Only Whitespace",
			input:    "     ",
			expected: []string{},
		},
		{
			name:     "Single Word No Spaces",
			input:    "singleword",
			expected: []string{"singleword"},
		},
		{
			name:     "Mixed Whitespace (Tabs and Newlines)",
			input:    "\tapple \n orange \r\n banana\t",
			expected: []string{"apple", "orange", "banana"},
		},
		{
			name:     "Symbols and Numbers",
			input:    "123 !@# $%^",
			expected: []string{"123", "!@#", "$%^"},
		},
		{
			name:     "Unicode and Emojis",
			input:    "Go is 🚀 awesome",
			expected: []string{"Go", "is", "🚀", "awesome"},
		},
		{
			name:     "Multi-byte Characters (Japanese)",
			input:    "世界 means world",
			expected: []string{"世界", "means", "world"},
		},
		{
			name:     "Multiple Consecutive Spaces",
			input:    "multiple         spaces",
			expected: []string{"multiple", "spaces"},
		},
	}
	passed := 0
	total := len(cases)
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			t.Logf("----------------")
			t.Logf("%sINPUT:%s \t\t %s", ColorYellow, ColorReset, c.input)
			actual := cleanInput(c.input)

			if len(actual) != len(c.expected) {
				t.Errorf("%s Lenght doesn't match expected %d got %d", FailureIcon, len(c.expected), len(actual))
				t.FailNow()
			}
			t.Logf("%s Length of EXPECTED %d and RESULT %d %sMATCH%s", SuccessIcon, len(c.expected), len(actual), ColorGreen, ColorReset)
			for i := range actual {
				word := actual[i]
				expectedWord := c.expected[i]
				if word != expectedWord {
					t.Errorf("%s Not right result expected %s got %s", FailureIcon, expectedWord, word)
					t.Fail()
				}
			}
			t.Logf("%s%s OUTPUT:%s \t\t %v", SuccessIcon, ColorGreen, ColorReset, actual)
			t.Logf("%s%s EXPECTING:%s \t %v", SuccessIcon, ColorGreen, ColorReset, c.expected)
			passed++
		})

	}
	t.Logf("--------------------")
	t.Logf("%d out of %d passed", passed, total)
	t.Logf("--------------------")
}
