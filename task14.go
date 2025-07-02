package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Простые числа до 100:")

	count := 0
	for num := 2; num <= 100; num++ {
		if isPrime(num) {
			fmt.Printf("%3d", num)
			count++
			if count%10 == 0 {
				fmt.Println()
			}
		}
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	maxDivisor := int(math.Sqrt(float64(n))) + 1
	for i := 3; i < maxDivisor; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
