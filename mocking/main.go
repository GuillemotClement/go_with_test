package main

import (
	"bytes"
	"fmt"
)

func main() {
	//Countdown()
}

// on passe en parametre de la fonction un pointeur vers le buffer qui sera utiliser pour stocker ce qu'on ecrit
func Countdown(out *bytes.Buffer) {
	fmt.Fprint(out, "3")
}
