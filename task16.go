package main

import (
	"fmt"
)

func main() {
	var dayNumber int

	fmt.Print("Введите номер дня недели (1-7): ")
	_, err := fmt.Scan(&dayNumber)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	dayName := getDayName(dayNumber)
	fmt.Printf("День недели №%d - это %s\n", dayNumber, dayName)
}

func getDayName(dayNumber int) string {
	switch dayNumber {
	case 1:
		return "Понедельник"
	case 2:
		return "Вторник"
	case 3:
		return "Среда"
	case 4:
		return "Четверг"
	case 5:
		return "Пятница"
	case 6:
		return "Суббота"
	case 7:
		return "Воскресенье"
	default:
		return "Неизвестный день (введите число от 1 до 7)"
	}
}
