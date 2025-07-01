package main

import (
	"fmt"
	"strings"
)

func main() {
	example := "Солнце светит круглый год вот"
	wordsAlt := strings.Split(example, " ")
	for _, word := range wordsAlt {
		fmt.Println(word)
	}
}
