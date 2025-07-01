package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "Солнце светит круглый год вот"
	charCount := len([]rune(input))

	fmt.Printf("1. Количество символов в строке: %d\n", charCount)

	substring := "светит"
	index := strings.Index(input, substring)
	if index != -1 {
		fmt.Printf("Подстрока '%s' найдена на позиции %d\n", substring, index)
	} else {
		fmt.Printf("Подстрока '%s' не найдена\n", substring)
	}

	lower := strings.ToLower(input)
	upper := strings.ToUpper(input)
	fmt.Println("3. Изменение регистра:")
	fmt.Println("   В нижнем регистре:", lower)
	fmt.Println("   В верхнем регистре:", upper)
}
