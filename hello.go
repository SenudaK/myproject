package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"

	espHelloPrefix = "Hola, "
	engHelloPrefix = "Hello, "
	fraHelloPrefix = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = espHelloPrefix
	case french:
		prefix = fraHelloPrefix
	default:
		prefix = engHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
