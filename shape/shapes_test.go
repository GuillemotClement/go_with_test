package shape

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.00, 10.00}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	// on definis une table driven test qui permet de passer une liste de test qui sera tester de la meme maniere
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		// on definis les test a realiser dans une struct anonyme
		// on peut ajouter le nom pour plus de lisibiliter avant de passer la valeur
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// tt.name depuis le case permet d'affiche le nom de la struct
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// %#v pour afficher la struct qui merde
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

	// TEST SANS LA TABLE DE TEST
	//checkArea := func(t testing.TB, shape Shape, want float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != want {
	//		t.Errorf("got %g want %g", got, want)
	//	}
	//}

	//t.Run("rectangles", func(t *testing.T) {
	//	rectangle := Rectangle{12, 6}
	//	checkArea(t, rectangle, 72.0)
	//})
	//
	//t.Run("circle", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkArea(t, circle, 314.1592653589793)
	//})

}
