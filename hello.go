package main

import "fmt"

func main() {
	fmt.Println(Hello("", ""))
}

const englishHelloPrefix = "Hello, "
const spannish = "Spanish"
const spannishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Salut, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spannish {
		return spannishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
