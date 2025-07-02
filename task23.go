package main

import (
	"fmt"
)

func LongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	fmt.Println(LongestString("яблоко", "банан", "абрикос"))
	fmt.Println(LongestString("один", "два", "десять", "four"))
	fmt.Println(LongestString("короткий", "длинный", "гигантский"))
	fmt.Println(LongestString(""))
	fmt.Println(LongestString("арбуз"))

	words := []string{"cat", "elephant", "dog", "giraffe"}
	fmt.Println(LongestString(words...))
}
