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
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		// --- Edge Cases ---
		{
			input:    "", // Empty string
			expected: []string{},
		},
		{
			input:    "     ", // Only whitespace
			expected: []string{},
		},
		{
			input:    "singleword", // No spaces
			expected: []string{"singleword"},
		},
		// --- Whitespace Varieties ---
		{
			input:    "\tapple \n orange \r\n banana\t", // Tabs and newlines
			expected: []string{"apple", "orange", "banana"},
		},
		// --- Special Characters ---
		{
			input:    "123 !@# $%^", // Numbers and symbols
			expected: []string{"123", "!@#", "$%^"},
		},
		{
			input:    "Go is 🚀 awesome", // Unicode/Emojis
			expected: []string{"Go", "is", "🚀", "awesome"},
		},
		{
			input:    "世界 means world", // Multi-byte characters
			expected: []string{"世界", "means", "world"},
		},
		// --- Formatting Stress Tests ---
		{
			input:    "multiple          spaces", // Large gaps
			expected: []string{"multiple", "spaces"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Lenght doesn't match expected %d got %d", len(c.expected), len(actual))
			t.Fail()
		}
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Not right result expected %s got %s", expectedWord, word)
				t.Fail()
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
	}
}
