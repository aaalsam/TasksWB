package main

import (
	"fmt"
	"strings"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// Могут возникнуть проблемы в виде не правильного копирования строки, потери символов и тд

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

//Строка - это слайс из байт. Это значит, что нарезая строку таким образом, мы получаем 100 байт. Однако, один символ может состоять из нескольких байт, которые мы могли просто не записать
// Из-за недостатка байт символы часто интерпретируются как � Чтобы исправить ошибку, мы должны нарезать слайс иначе

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
	//Нарезав строку как руны (int32 коды символов стандарта Unicode), мы правильно определим байтовые границы каждого символа
	runes := []rune(v)
	justString = string(runes[:100])
	fmt.Println(justString)
}

func main() {
	someFunc()
}
