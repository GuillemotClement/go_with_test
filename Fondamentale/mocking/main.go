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

type SpyCountdwonOperations struct {
	Calls []string
}

func (s *SpyCountdwonOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdwonOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const write = "write"
const sleep = "sleep"

// definition des constante
const finalWorld = "Go!"
const countDownStart = 3

// V2
// on rend la fonction generique en lui passant une interface io.Writter
// cela permet de rendre la fonction generique et lui permettant d'accepter n'importe quel objet qui implement l'interface (pour ecriture)
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
	}
	for i := countDownStart; i > 0; i-- {
		fmt.Println(out, i)
	}

	fmt.Fprint(out, finalWorld)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	// os.Stdout permet d'ecrire dans le terminal (la sortie standard)
	Countdown(os.Stdout, sleeper)
}
