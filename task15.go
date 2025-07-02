package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&a)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err)
		return
	}

	fmt.Print("Введите оператор (+, -, *, /): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Ошибка ввода оператора:", err)
		return
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanln(&b)
	if err != nil {
		fmt.Println("Ошибка ввода числа:", err)
		return
	}

	result, err := calculate(a, b, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
		os.Exit(1)
	}

	fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", a, operator, b, result)
}

func calculate(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("деление на ноль")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("неподдерживаемая операция: %s", operator)
	}
}
