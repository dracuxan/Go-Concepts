package main

import "strings"

const (
	helloStringEnglish = "Hello, "
	helloStringSpanish = "Hola, "
	helloStringFrench  = "Bonjour, "
	spanish            = "spanish"
	french             = "french"
)

func helloWorld(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return setPrrefix(language) + name
}

func setPrrefix(language string) (prefix string) {
	switch strings.ToLower(language) {
	case spanish:
		prefix = helloStringSpanish
	case french:
		prefix = helloStringFrench
	default:
		prefix = helloStringEnglish
	}
	return
}
