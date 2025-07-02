package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Name() string
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Name() string {
	return "Прямоугольник"
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Name() string {
	return "Круг"
}

func PrintArea(s Shape) {
	fmt.Printf("%s: площадь = %.2f\n", s.Name(), s.Area())
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}
	shapes := []Shape{rect, circle}

	for _, shape := range shapes {
		PrintArea(shape)
	}

	smallCircle := Circle{Radius: 2.5}
	bigRectangle := Rectangle{Width: 15, Height: 20}

	PrintArea(smallCircle)
	PrintArea(bigRectangle)
}
