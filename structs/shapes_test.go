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
	t.Run("for a rectangle", func(t *testing.T) {
		rectangle := Rectangle{3.0, 6.0}
		got := rectangle.Area()
		want := 18.0

		if got != want {
			t.Errorf("Expected %g, but got %g", want, got)
		}
	})

	t.Run("for a circle", func(t *testing.T) {
		circle := Circle{4.5}
		got := circle.Area()
		want := 28.274333882308138

		if got != want {
			t.Errorf("Expected %g, but got %g", want, got)
		}
	})
}
