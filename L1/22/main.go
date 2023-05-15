package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значение которых > 2^20.

// Программа выйдет за пределы допустимых значений, что может привести к неожиданным результатам. Максимально положительно  uint 18 446 744 073 709 551 615, отрицательное –9 223 372 036 854 775 808
// big.Int для крупных целых чисел, когда 18 квинтиллионов будет недостаточно.

func sum(a, b *big.Int) *big.Int {
	var c big.Int // Создаем переменную типа big.Int
	c.Add(a, b) // Вызываем метод у это переменной
	return &c
}

func dif(a, b *big.Int) *big.Int {
	var c big.Int
	c.Neg(b)
	return sum(a, &c)
}

func mul(a, b *big.Int) *big.Int {
	var c big.Int
	c.Mul(a, b)
	return &c
}

func div(a, b *big.Int) *big.Int {
	var c big.Int
	c.Quo(a, b)
	return &c
}

func main() {

	bigIntA := big.NewInt(int64(1 << 30))
	bigIntB := big.NewInt(int64(1 << 25))

	fmt.Println(bigIntA)
	fmt.Println(bigIntB)

	fmt.Println(sum(bigIntA, bigIntB))

	fmt.Println(dif(bigIntA, bigIntB))

	fmt.Println(mul(bigIntA, bigIntB))

	fmt.Println(div(bigIntA, bigIntB))

}