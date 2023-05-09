package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

var justString string

func createHugeString(n int) string {

	hugeString := strings.Builder{}
	for i := 0; i < n; i++ {
		hugeString.WriteString("a")
	}
	return hugeString.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	//justString = v[:100]

	tmp := make([]rune, 100)
	copy(tmp, []rune(v)[:100])
	justString = string(tmp) // конвертация в строки

	fmt.Println("justString = ", justString)
	fmt.Println("v = ", v)
}

func main() {
	someFunc()
}
