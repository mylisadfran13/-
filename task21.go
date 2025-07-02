package main

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
