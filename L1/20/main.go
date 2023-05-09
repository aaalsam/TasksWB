package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».

func main() {
	s := "snow dog sun"
	fmt.Println(s)
	fmt.Println(turnString(s))
}

func turnString(s string) string {
	w := strings.Fields(s) // snow, dog, sun
	n := len(w)
	for i := 0; i < n/2; i++ { // 
		w[i], w[n-i-1] = w[n-i-1], w[i]
	}
	return strings.Join(w, " ")
}
