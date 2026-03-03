package main

import (
	"strings"
)

func cleanInput(text string) []string {
	cleanText := strings.TrimSpace(text)
	splitText := strings.Split(cleanText, " ")
	return splitText
}
