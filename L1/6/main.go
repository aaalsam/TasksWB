package main

import (
	"context"
	"fmt"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.

func main() {

	ctx, cancel := context.WithCancel(context.Background()) //  Закрытие с помощью контекста с отменой

	ch := make(chan int)

	go func1(ch, ctx)

	ch <- 1

	cancel() // Отмена контекста 

	time.Sleep(1 * time.Second) 

}

func func1(ch chan int, ctx context.Context) {
	for {
		select {
		case a := <-ch:
			fmt.Println(a)
		case <-ctx.Done(): // Закрытие канала от имени контекста
			return
		}
	}
}
