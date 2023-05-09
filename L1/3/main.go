package main

import "fmt"

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

func main() {
	var numbers = []int{2, 4, 6, 8, 10}

	ch := make(chan int)
	for _, num := range numbers {
		go numberSquared(num, ch) // создаётся отдельная горутина, добавляющая к приложению другой поток Go, отвечающий за выполнение функции на отдельном конкурентном потоке
	}

	sum := 0
	for range numbers {
		sum += <-ch // инкрементирует на значение, считанное из канала
	}

	fmt.Println(sum)
}


func numberSquared(num int, ch chan int) { // запись квадрата числа в канал
	ch <- num * num
}
