package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		pi = math.Pi
		e  = math.E
	)

	radius := 5.0
	circleArea := pi * math.Pow(radius, 2)
	fmt.Println("Площадь кргуа с радиусом", radius, "равна:", circleArea)

	x := 2.0
	exponential := math.Pow(e, x)
	fmt.Println("e^x = ", exponential)
}
