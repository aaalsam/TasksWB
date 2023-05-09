package main

import (
	"fmt"
	"sync"
	"time"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

var wg sync.WaitGroup

func main() {

	array := [5]int{2, 4, 6, 8, 10}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go write1(array, ch1)
	go write2(ch1, ch2)
	go read(ch2)

	time.Sleep(1 * time.Millisecond)

	wg.Wait() // Читающие горутины завершат работу после закрытия каналов пишущими горутинами
}

func write1(array [5]int, ch1 chan int) { // Запись в первый канал значений х из массива
	wg.Add(1)
	for _, num := range array {
		ch1 <- num
	}
	close(ch1)
	wg.Done()
}

func write2(ch1 chan int, ch2 chan int) { // Запись во второй канал значений из первого канала * 2
	wg.Add(1)
	for num := range ch1 {
		ch2 <- num * 2
	}
	close(ch2)
	wg.Done()
}

func read(ch2 chan int) { // Чтение из c2 и запись в os.Stdout
	wg.Add(1)
	for num := range ch2 {
		fmt.Println(num)
	}
	wg.Done()
}
