package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

func uniqueness(s string) bool {

	set := make(map[rune]struct{})

	letters := strings.ToLower(s) // Приведение символа к нижнему регистру

	for _, l := range letters {
		if _, ok := set[l]; ok { // Если такой символ уже есть во множестве, возвращаем false
			return false
		}

		set[l] = struct{}{} // Если символ ранее не встречался, добавляем его во множество
	}
	return true
}

func main() {
	str := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, i := range str {
		fmt.Println(uniqueness(i))
	}
}
