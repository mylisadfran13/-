package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func modifyByValue(p Person) {
	p.Name = "Имя изменено (по значению)"
	p.Age = 100
	fmt.Println("Внутри modifyByValue:", p)
}

func modifyByPointer(p *Person) {
	p.Name = "Имя изменено (по указателю)"
	p.Age = 200
	fmt.Println("Внутри modifyByPointer:", *p)
}

func modifyIntByValue(num int) {
	num = 100
	fmt.Println("Внутри modifyIntByValue:", num)
}

func modifyIntByPointer(num *int) {
	*num = 200
	fmt.Println("Внутри modifyIntByPointer:", *num)
}

func main() {
	person := Person{Name: "Иван", Age: 30}

	fmt.Println("\n=== Работа со структурой ===")
	fmt.Println("Оригинальное значение:", person)

	modifyByValue(person)
	fmt.Println("После modifyByValue:", person)

	modifyByPointer(&person)
	fmt.Println("После modifyByPointer:", person)

	num := 1

	fmt.Println("\n=== Работа с простым типом ===")
	fmt.Println("Оригинальное значение:", num)

	modifyIntByValue(num)
	fmt.Println("После modifyIntByValue:", num)

	modifyIntByPointer(&num)
	fmt.Println("После modifyIntByPointer:", num)

	fmt.Println("\n=== Сравнение производительности ===")
	bigData := make([]int, 1000000)

	fmt.Println("Передача большого среза по значению (создается копия)")
	processByValue(bigData)

	fmt.Println("Передача большого среза по указателю (копирования нет)")
	processByPointer(&bigData)
}

func processByValue(data []int) {
	_ = len(data)
}

func processByPointer(data *[]int) {
	_ = len(*data)
}
