package main

import "fmt"

func main() {
	fmt.Println(Hello("world", ""))
}

const spanish = "Spanish"
const spanishPrefix = "Hola, "

const french = "French"
const frenchPrefix = "Bonjour, "

const englishHelloPrefix = "Hello, "

func Hello(user, language string) string {
	if user == "" {
		user = "World"
	}
	return greetingPrefix(language) + user + "!"

}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = spanishPrefix
	case french:
		prefix = frenchPrefix
	default:
		prefix = englishHelloPrefix

	}

	return
}
