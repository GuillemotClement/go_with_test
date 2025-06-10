package main

import "fmt"

func main() {
	fmt.Println(Hello("", ""))
}

// on prepare les differente constante utiliser dans la fonction
// on peut declarer un groupe de constante
const (
	spannish = "Spanish"
	french   = "French"

	englishHelloPrefix  = "Hello, "
	spannishHelloPrefix = "Hola, "
	frenchHelloPrefix   = "Salut, "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// on ajoute un retour nommer qui permet de declare automatiqement la variable dans la fonction
// prend automatiquement la valeur nul lier au type => ici ""
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spannish:
		prefix = spannishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return // on retourne le prefix qui est declarer dans le retour de la fonction
}
