package main

import (
	"fmt"
)

func RemoveElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func main() {
	fruits := []string{"Яблоко", "Банан", "Апельсин", "Груша", "Киви"}
	fmt.Println("Исходный срез:", fruits)

	fruits = RemoveElement(fruits, 2)
	fmt.Println("После удаления элемента с индексом 2:", fruits)

	fruits = RemoveElement(fruits, 0)
	fmt.Println("После удаления первого элемента:", fruits)

	fruits = RemoveElement(fruits, 10)
	fmt.Println("Попытка удаления несуществующего элемента:", fruits)
}
