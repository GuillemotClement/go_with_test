package main

import (
	"fmt"
	"io"
	"os"
)

// V1
// on passe en parametre de la fonction un pointeur vers le buffer qui sera utiliser pour stocker ce qu'on ecrit
//func Countdown(out *bytes.Buffer) {
//	fmt.Fprint(out, "3")
//}

// V2
// on rend la fonction generique en lui passant une interface io.Writter
// cela permet de rendre la fonction generique et lui permettant d'accepter n'importe quel objet qui implement l'interface (pour ecriture)
func Countdown(out io.Writer) {
	fmt.Fprint(out, "3")
}

func main() {
	// os.Stdout permet d'ecrire dans le terminal (la sortie standard)
	Countdown(os.Stdout)
}
