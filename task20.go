package main

import (
	"fmt"
	"strings"
)

type Engine struct {
	Type         string
	Volume       float64
	Power        int
	Manufacturer string
}

type Car struct {
	Brand     string
	Model     string
	Year      int
	Mileage   int
	Engine    Engine
	IsRunning bool
}

func NewCar(brand, model string, year, mileage int, engine Engine) (*Car, error) {
	if strings.TrimSpace(brand) == "" {
		return nil, fmt.Errorf("марка не может быть пустой")
	}
	if strings.TrimSpace(model) == "" {
		return nil, fmt.Errorf("модель не может быть пустой")
	}
	if year < 1886 || year > 2023 {
		return nil, fmt.Errorf("некорректный год выпуска")
	}
	if mileage < 0 {
		return nil, fmt.Errorf("пробег не может быть отрицательным")
	}

	return &Car{
		Brand:     strings.TrimSpace(brand),
		Model:     strings.TrimSpace(model),
		Year:      year,
		Mileage:   mileage,
		Engine:    engine,
		IsRunning: false,
	}, nil
}

func (c *Car) Start() {
	if !c.IsRunning {
		c.IsRunning = true
		fmt.Printf("%s %s заведен\n", c.Brand, c.Model)
	} else {
		fmt.Printf("%s %s уже работает\n", c.Brand, c.Model)
	}
}

func (c *Car) Stop() {
	if c.IsRunning {
		c.IsRunning = false
		fmt.Printf("%s %s заглушен\n", c.Brand, c.Model)
	} else {
		fmt.Printf("%s %s уже не работает\n", c.Brand, c.Model)
	}
}

func (c *Car) Drive(distance int) error {
	if !c.IsRunning {
		return fmt.Errorf("невозможно ехать - двигатель не работает")
	}
	if distance <= 0 {
		return fmt.Errorf("дистанция должна быть положительной")
	}

	c.Mileage += distance
	fmt.Printf("%s %s проехал %d км. Общий пробег: %d км\n",
		c.Brand, c.Model, distance, c.Mileage)
	return nil
}

func (c *Car) GetInfo() string {
	return fmt.Sprintf(
		"Автомобиль: %s %s\n"+
			"Год выпуска: %d\n"+
			"Пробег: %d км\n"+
			"Состояние: %v\n"+
			"Двигатель: %s %.1f л (%d л.с.) от %s",
		c.Brand, c.Model,
		c.Year,
		c.Mileage,
		c.IsRunning,
		c.Engine.Type, c.Engine.Volume, c.Engine.Power, c.Engine.Manufacturer,
	)
}

func main() {
	engine := Engine{
		Type:         "бензиновый",
		Volume:       2.0,
		Power:        150,
		Manufacturer: "Toyota",
	}

	car, err := NewCar("Toyota", "Camry", 2020, 50000, engine)
	if err != nil {
		fmt.Println("Ошибка создания автомобиля:", err)
		return
	}

	fmt.Println(car.GetInfo())
	fmt.Println(" ")

	car.Start()
	car.Start()
	fmt.Println(" ")

	err = car.Drive(150)
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	fmt.Println(" ")

	car.Stop()
	car.Stop()
	fmt.Println(" ")

	_, err = NewCar("", "Model S", 2025, -100, engine)
	if err != nil {
		fmt.Println("Ошибка создания автомобиля:", err)
	}
}
