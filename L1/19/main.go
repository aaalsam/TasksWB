package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func main(){
	s := "Hello"
	fmt.Println(reverse(s))
}

func reverse(s string) string {
    r := []rune(s) // Создаем слайс рун по строке
    for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}