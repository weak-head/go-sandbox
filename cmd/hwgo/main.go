package main

import "fmt"

const (
	spanish       = "Spanish"
	french        = "French"
	exclamation   = "!"
	spanishPrefix = "Hola, "
	englishPrefix = "hwgo, "
	frenchPrefix  = "Bonjour, "
)

// Hello returns greeting message
func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name + exclamation
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("GO", ""))
}
