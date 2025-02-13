package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(2.0, 4.0)
	want := 12.0

	if got != want {
		t.Errorf("Expected %.2f, but got %.2f", want, got)
	}
}

func TestArea(t *testing.T) {
	got := Area(3.0, 6.0)
	want := 18.0

	if got != want {
		t.Errorf("Expected %.2f, but got %.2f", want, got)
	}
}
