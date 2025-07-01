package main

import "fmt"

func main() {
	var slice []string

	slice = append(slice, "первый")
	slice = append(slice, "второй")
	slice = append(slice, "третий")

	fmt.Println("Все элементы среза:")
	for i, item := range slice {
		fmt.Printf("Индекс: %d, Значение: %s\n", i, item)
	}
}
