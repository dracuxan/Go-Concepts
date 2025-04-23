package helloworld

import "strings"

const (
	spanish = "spanish"
	french  = "french"

	helloStringEnglish = "Hello, "
	helloStringSpanish = "Hola, "
	helloStringFrench  = "Bonjour, "
)

func HelloWorld(name string, language string) string {
	if name == "" {
		name = "Unknown"
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
