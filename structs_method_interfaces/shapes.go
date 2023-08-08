package structsmethodinterfaces

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Height + 2*r.Width
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Circle struct {
	Radius float64
}
