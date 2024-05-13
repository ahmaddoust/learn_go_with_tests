package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	pashto  = "Pashto"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	pashtoHelloPrefix  = "Salam, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// prefix := englishHelloPrefix

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case pashto:
		prefix = pashtoHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// 	if language == spanish {
// 		return spanishHelloPrefix + name
// 	}
// 	if language == french {
// 		return frenchHelloPrefix + name
// 	}
// 	return englishHelloPrefix + name
// }

func main() {
	fmt.Println(Hello("world", ""))
}
