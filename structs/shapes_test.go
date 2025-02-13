package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(2.0, 4.0)
	want := 10.0

	if got != want {
		t.Errorf("Expected %.2f, but got %.2f", want, got)
	}
}
