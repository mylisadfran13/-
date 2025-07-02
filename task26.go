package main

import (
	"fmt"
	"strings"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Grades    []int
}

func NewStudent(name string, birthYear int, grades []int) (*Student, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("имя не может быть пустым")
	}
	if birthYear < 1900 || birthYear > time.Now().Year() {
		return nil, fmt.Errorf("некорректный год рождения")
	}
	for _, grade := range grades {
		if grade < 1 || grade > 5 {
			return nil, fmt.Errorf("оценка должна быть от 1 до 5")
		}
	}

	return &Student{
		Name:      strings.TrimSpace(name),
		BirthYear: birthYear,
		Grades:    grades,
	}, nil
}

func (s *Student) CalculateAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s *Student) GetStatus() string {
	if len(s.Grades) == 0 {
		return "нет оценок"
	}

	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	average := float64(sum) / float64(len(s.Grades))

	switch {
	case average >= 4.5:
		return "отличник"
	case average >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func (s *Student) AddGrade(grade int) error {
	if grade < 1 || grade > 5 {
		return fmt.Errorf("оценка должна быть от 1 до 5")
	}
	s.Grades = append(s.Grades, grade)
	return nil
}

func (s *Student) GetInfo() string {
	return fmt.Sprintf(
		"Студент: %s\nВозраст: %d\nСтатус: %s\nОценки: %v",
		s.Name,
		s.CalculateAge(),
		s.GetStatus(),
		s.Grades,
	)
}

func main() {
	student1, _ := NewStudent("Александр В", 19, []int{5, 5, 4, 5})
	student2, _ := NewStudent("Андрей П", 18, []int{3, 4, 3, 4})
	student3, _ := NewStudent("Юлия К", 18, []int{3, 3, 3, 2})

	student1.AddGrade(5)
	student2.AddGrade(5)
	student3.AddGrade(4)

	fmt.Println(student1.GetInfo())
	fmt.Println("\n" + strings.Repeat("-", 30))
	fmt.Println(student2.GetInfo())
	fmt.Println("\n" + strings.Repeat("-", 30))
	fmt.Println(student3.GetInfo())

	student4, err := NewStudent("", 1800, []int{6})
	if err != nil {
		fmt.Println("\nОшибка создания студента:", err)
	}
}
