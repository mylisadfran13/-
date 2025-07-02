package main

import "fmt"

func main() {
	fmt.Println("Таблица умножения от 1 до 10")
	fmt.Println("============================")

	fmt.Printf("%5s", "x")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println("\n" + "     " + "-----" + "-----" + "-----" + "-----" + "-----" +
		"-----" + "-----" + "-----" + "-----" + "-----")

	for i := 1; i <= 10; i++ {
		fmt.Printf("%5d", i)
		for j := 1; j <= 10; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}
