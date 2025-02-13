package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{2.0, 4.0}
	got := Perimeter(rectangle)
	want := 12.0

	if got != want {
		t.Errorf("Expected %.2f, but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{3.0, 6.0}
	got := Area(rectangle)
	want := 18.0

	if got != want {
		t.Errorf("Expected %.2f, but got %.2f", want, got)
	}
}
