package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Ошибка ввода первого числа:", err)
		return
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Ошибка ввода второго числа:", err)
		return
	}

	fmt.Printf("\nРезультаты операций с числами %.2f и %.2f:\n", a, b)
	fmt.Println(" ")

	fmt.Printf("%.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("%.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("%.2f * %.2f = %.2f\n", a, b, a*b)

	if b != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Printf("%.2f / %.2f = невозможно (деление на ноль)\n", a, b)
	}

	if b != 0 {
		fmt.Printf("%d %% %d = %d\n", int(a), int(b), int(a)%int(b))
	} else {
		fmt.Printf("%d %% %d = невозможно (деление на ноль)\n", int(a), int(b))
	}

	fmt.Printf("%.2f ^ %.2f = %.2f\n", a, b, math.Pow(a, b))

	if b != 0 {
		fmt.Printf("%d // %d = %d\n", int(a), int(b), int(a)/int(b))
	} else {
		fmt.Printf("%d // %d = невозможно (деление на ноль)\n", int(a), int(b))
	}
}
