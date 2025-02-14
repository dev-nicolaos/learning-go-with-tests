package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius  float64
}

func (rectangle *Rectangle) Perimeter() float64 {
	return rectangle.Width*2 + rectangle.Height*2
}

func (rectangle *Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

func (circle *Circle) Area() float64 {
	return 2*circle.Radius*math.Pi
}
