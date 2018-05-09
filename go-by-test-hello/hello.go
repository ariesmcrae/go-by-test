package main

import "fmt"

const french = "French"
const spanish = "Spanish"

const helloPrefixEnglish = "Hello, "
const helloPrefixFrench = "Bonjour, "
const helloPrefixSpanish = "Hola, "

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = helloPrefixFrench
	case spanish:
		prefix = helloPrefixSpanish
	default:
		prefix = helloPrefixEnglish
	}

	return
}
