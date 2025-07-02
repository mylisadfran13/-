package main

import (
	"fmt"
)

type Transport interface {
	Move()
	Stop()
	Name() string
}

type Car struct {
	Model string
	Speed int
}

type Bicycle struct {
	Brand string
	Gears int
}

type Airplane struct {
	Airliner string
	Altitude int
}

func (c Car) Move() {
	fmt.Printf("Автомобиль %s едет со скоростью %d км/ч\n", c.Model, c.Speed)
}

func (c Car) Stop() {
	fmt.Printf("Автомобиль %s тормозит\n", c.Model)
}

func (c Car) Name() string {
	return fmt.Sprintf("Автомобиль %s", c.Model)
}

func (b Bicycle) Move() {
	fmt.Printf("Велосипед %s крутит педали (передача %d)\n", b.Brand, b.Gears)
}

func (b Bicycle) Stop() {
	fmt.Printf("Велосипед %s останавливается\n", b.Brand)
}

func (b Bicycle) Name() string {
	return fmt.Sprintf("Велосипед %s", b.Brand)
}

func (a Airplane) Move() {
	fmt.Printf("Самолет %s взлетает на высоту %d метров\n", a.Airliner, a.Altitude)
}

func (a Airplane) Stop() {
	fmt.Printf("Самолет %s приземляется\n", a.Airliner)
}

func (a Airplane) Name() string {
	return fmt.Sprintf("Самолет %s", a.Airliner)
}

func TestTransport(t Transport) {
	fmt.Println("\nТестируем:", t.Name())
	t.Move()
	t.Stop()
}

func main() {
	car := Car{Model: "Toyota Camry", Speed: 60}
	bike := Bicycle{Brand: "Trek", Gears: 7}
	plane := Airplane{Airliner: "Boeing 747", Altitude: 10000}

	transports := []Transport{car, bike, plane}

	for _, t := range transports {
		TestTransport(t)
	}

	sportCar := Car{Model: "Ferrari", Speed: 120}
	mountainBike := Bicycle{Brand: "Giant", Gears: 21}

	TestTransport(sportCar)
	TestTransport(mountainBike)
}

//go run hello.go
