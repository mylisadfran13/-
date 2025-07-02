package main

import (
	"fmt"
	"sort"
)

// базовые функции
func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}
	return -1
}

func Remove[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}

func Unique[T comparable](slice []T) []T {
	seen := make(map[T]bool)
	result := []T{}
	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}

// сортировка
func Sort[T ~string | ~int | ~float64](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
}

func SortBy[T any](slice []T, less func(i, j int) bool) {
	sort.Slice(slice, less)
}

func SortDesc[T ~string | ~int | ~float64](slice []T) {
	sort.Slice(slice, func(i, j int) bool {
		return slice[i] > slice[j]
	})
}

// фильтрация
func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0)
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T, U any](slice []T, mapper func(T) U) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = mapper(item)
	}
	return result
}

func Reduce[T, U any](slice []T, initial U, reducer func(U, T) U) U {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

func main() {
	numbers := []int{5, 2, 8, 3, 1, 9, 4, 2, 5}
	words := []string{"Яблоко", "Банан", "Апельсин", "Яблоко", "Киви"}

	fmt.Println("Contains 8:", Contains(numbers, 8))
	fmt.Println("Index of 'Апельсин':", IndexOf(words, "Апельсин"))

	uniqueNumbers := Unique(numbers)
	fmt.Println("Unique numbers:", uniqueNumbers)

	Sort(numbers)
	fmt.Println("Sorted numbers:", numbers)

	SortDesc(words)
	fmt.Println("Sorted words desc:", words)

	evenNumbers := Filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("Even numbers:", evenNumbers)

	lengths := Map(words, func(s string) int {
		return len(s)
	})
	fmt.Println("Word lengths:", lengths)

	total := Reduce(numbers, 0, func(sum, n int) int {
		return sum + n
	})
	fmt.Println("Sum of numbers:", total)

	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{"Юля", 25},
		{"Саша", 30},
		{"Андрей", 20},
	}

	SortBy(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted people by age:", people)
}
