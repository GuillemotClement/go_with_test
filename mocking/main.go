package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// V1
// on passe en parametre de la fonction un pointeur vers le buffer qui sera utiliser pour stocker ce qu'on ecrit
//
//	func Countdown(out *bytes.Buffer) {
//		fmt.Fprint(out, "3")
//	}
type Sleeper interface {
	Sleep()
}

// permet de connaitre le nombre d'appel de la fonction sleep()
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// definition des constante
const finalWorld = "Go!"
const countDownStart = 3

// V2
// on rend la fonction generique en lui passant une interface io.Writter
// cela permet de rendre la fonction generique et lui permettant d'accepter n'importe quel objet qui implement l'interface (pour ecriture)
func Countdown(out io.Writer, sleeper Sleeper) {
	// boucle qui permet de faire le decompte
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		// permet d'ajoute un temps entre chaque iteration
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWorld)
}

func main() {
	sleeper := &DefaultSleeper{}
	// os.Stdout permet d'ecrire dans le terminal (la sortie standard)
	Countdown(os.Stdout, sleeper)
}
