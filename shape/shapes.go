package shape

import "math"

// definition d'une interface.
// on viens definir des methodes avec le retour dedans
// une interface est automatiquement implementer par une struct si elle implemente toutes les methodes de l'interface
type Shape interface {
	Area() float64
}
type Rectangle struct {
	Width  float64
	Height float64
}

// on passe la struct en parametre
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

// on ajoute la methode a la struct
// la methode a acces aux prop de la struct comme pour une classe
// pour lier une methode a une struct on passe un receiver avant le nom de la methode => par convention premiere lettre de la struct
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
