package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.0, 4.0}
	got := rectangle.Perimeter()
	want := 12.0

	if got != want {
		t.Errorf("Expected %.2f, but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	checkArea := func (t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("Expected %g, but got %g", want, got)
		}
	}

	t.Run("for a rectangle", func(t *testing.T) {
		checkArea(t, Rectangle{3.0, 6.0}, 18.0)
	})

	t.Run("for a circle", func(t *testing.T) {
		checkArea(t, Circle{4.5}, 28.274333882308138)
	})
}
