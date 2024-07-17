package structuresMethodsAndInterfaces

import "math"

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}
type Triangle struct {
	Base, Height float64
}
type Form interface {
	Area() float64
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
