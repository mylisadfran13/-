package main

import (
	"fmt"
	"strings"
	"unicode"
)

type StudentGrades map[string]int

func (sg StudentGrades) AddGrade(name string, grade int) {
	name = normalizeName(name)
	sg[name] = grade
	fmt.Printf("Оценка для студента %s: %d\n", name, grade)
}

func (sg StudentGrades) GetGrade(name string) (int, bool) {
	name = normalizeName(name)
	grade, exists := sg[name]
	return grade, exists
}

func (sg StudentGrades) RemoveGrade(name string) {
	name = normalizeName(name)
	if _, exists := sg[name]; exists {
		delete(sg, name)
		fmt.Printf("Студент %s удален\n", name)
	} else {
		fmt.Printf("Студент %s не найден\n", name)
	}
}

func (sg StudentGrades) PrintAll() {
	fmt.Println("\nВсе оценки:")
	for name, grade := range sg {
		fmt.Printf("%s: %d\n", name, grade)
	}
}

func normalizeName(name string) string {
	name = strings.ToLower(name)
	return capitalizeFirstLetters(name)
}

func capitalizeFirstLetters(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		if len(word) > 0 {
			r := []rune(word)
			r[0] = unicode.ToUpper(r[0])
			words[i] = string(r)
		}
	}
	return strings.Join(words, " ")
}

func main() {
	grades := make(StudentGrades)

	grades.AddGrade("Александр В", 5)
	grades.AddGrade("Андрей П", 4)
	grades.AddGrade("Юлия К", 3)

	grades.AddGrade("юлия к", 4)

	if grade, exists := grades.GetGrade("Александр В"); exists {
		fmt.Printf("Оценка Александра В: %d\n", grade)
	}

	if _, exists := grades.GetGrade("Неизвестный Студент"); !exists {
		fmt.Println("Студент не найден")
	}

	grades.RemoveGrade("Андрей П")

	grades.RemoveGrade("Анна Каренина")

	grades.PrintAll()
}
