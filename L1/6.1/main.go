package main

import (
	"fmt"
	"sync"
	"time"
)

// Завершается естественным путем

func main() {

	wg := new(sync.WaitGroup)

	ch := make(chan int)

	go func1(ch, wg)
	
	time.Sleep(1 * time.Second)

	ch <- 1

	wg.Wait() // Завершается группа и с ней горутина, когда счетчик в группе станет 0
}

func func1(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1) // Добавляем в группу
	fmt.Println(<-ch)
	wg.Done() // Минус 1 из группы
}
