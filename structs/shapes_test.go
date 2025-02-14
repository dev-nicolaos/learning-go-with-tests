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
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"rectangele", Rectangle{3.0, 6.0}, 18.0},
		{"circle", Circle{4.5}, 28.274333882308138},
		{"triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("Expected %g, but got %g", tt.want, got)
			}
		})
	}
}
