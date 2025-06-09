package main

import "fmt"

func main() {
	fmt.Println(Hello("", ""))
}

// on prepare les differente constante utiliser dans la fonction
const englishHelloPrefix = "Hello, "
const spannish = "Spanish"
const spannishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Salut, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	// on prepare le prefix
	prefix := englishHelloPrefix
	// on utilise le switch pour tester les valeurs passer a la fonction
	switch language {
	case spannish:
		prefix = spannishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}

	return prefix + name
}
