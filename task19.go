package main

import (
	"fmt"
	"strings"
)

type Student struct {
	Name     string
	Age      int
	Course   int
	AvgGrade float64
}

func NewStudent(name string, age, course int, avgGrade float64) (*Student, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("имя не может быть пустым")
	}
	if age <= 0 {
		return nil, fmt.Errorf("возраст должен быть положительным числом")
	}
	if course < 1 || course > 6 {
		return nil, fmt.Errorf("курс должен быть от 1 до 6")
	}
	if avgGrade < 0 || avgGrade > 10 {
		return nil, fmt.Errorf("средний балл должен быть от 0 до 10")
	}

	return &Student{
		Name:     strings.TrimSpace(name),
		Age:      age,
		Course:   course,
		AvgGrade: avgGrade,
	}, nil
}

func (s *Student) Promote() error {
	if s.Course >= 6 {
		return fmt.Errorf("студент уже на последнем курсе")
	}
	s.Course++
	return nil
}

func (s *Student) UpdateGrade(newGrade float64) error {
	if newGrade < 0 || newGrade > 10 {
		return fmt.Errorf("оценка должна быть от 0 до 10")
	}
	s.AvgGrade = newGrade
	return nil
}

func (s *Student) GetInfo() string {
	return fmt.Sprintf(
		"Студент: %s\nВозраст: %d\nКурс: %d\nСредний балл: %.2f",
		s.Name, s.Age, s.Course, s.AvgGrade,
	)
}

func main() {

	student, err := NewStudent("Александр В", 20, 2, 8.5)
	if err != nil {
		fmt.Println("Ошибка создания студента:", err)
		return
	}

	fmt.Println("Информация о студенте:")
	fmt.Println(student.GetInfo())
	fmt.Println("---------------------")

	err = student.Promote()
	if err != nil {
		fmt.Println("Ошибка перевода на следующий курс:", err)
	} else {
		fmt.Println("Студент переведен на следующий курс!")
		fmt.Printf("Текущий курс: %d\n", student.Course)
	}
	fmt.Println(" ")

	err = student.UpdateGrade(9.1)
	if err != nil {
		fmt.Println("Ошибка обновления оценки:", err)
	} else {
		fmt.Println("Средний балл обновлен!")
		fmt.Printf("Новый средний балл: %.2f\n", student.AvgGrade)
	}
	fmt.Println(" ")

	_, err = NewStudent("", -1, 0, 11)
	if err != nil {
		fmt.Println("Ошибка создания студента:", err)
	}
}
