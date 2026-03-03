package main

import (
	"strings"
)

func cleanInput(text string) []string {
	splitText := strings.Fields(text)
	//This is the wrong way, but i keeping it here for looking at formating of unit test if failed
	//splitText := strings.Split(text, " ")
	return splitText
}
