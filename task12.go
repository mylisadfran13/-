package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := `«В четве́рг четвёртого числа́ в четы́ре с че́твертью часа́
	лигури́йский регулиро́вщик регули́ровал в Лигу́рии.
	Но три́дцать три корабля́ лави́ровали, лави́ровали,
	да так и не вы́лавировали`

	wordFrequency := countWords(text)

	fmt.Println("Частота слов в тексте:")
	for word, count := range wordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func countWords(text string) map[string]int {

	frequency := make(map[string]int)
	isLetter := func(r rune) bool {
		return unicode.IsLetter(r)
	}

	words := strings.FieldsFunc(text, func(r rune) bool {
		return !isLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		frequency[lowerWord]++
	}

	return frequency
}
