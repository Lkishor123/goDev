package interfacestruct

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
	fmt.Println("Area: ", s.Area())
}

func Interfacemain() {
	r := Rectangle{Width: 3, Height: 4}
	c := Circle{Radius: 5}
	PrintArea(r)
	PrintArea(c)
}
