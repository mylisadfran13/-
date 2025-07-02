package main

import "fmt"

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	x := 10
	y := 20

	fmt.Printf("До обмена: x = %d, y = %d\n", x, y)
	Swap(&x, &y)

	fmt.Printf("После обмена: x = %d, y = %d\n", x, y)

	a := 100
	b := 200
	fmt.Printf("\nДо обмена: a = %d, b = %d\n", a, b)
	Swap(&a, &b)
	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
