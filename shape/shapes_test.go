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
		shape Shape
		want  float64
	}{
		// on definis les test a realiser dans une struct anonyme
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
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
