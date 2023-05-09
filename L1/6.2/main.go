package main

import (
	"fmt"
	"time"
)

func main() {

	quit := make(chan int)

	go func1(quit)

	time.Sleep(1 * time.Second)

	quit <- 1 // Канал, в который отправляется сообщение, при этом горутина завершает работу 
}

func func1(quit chan int) {
	for {
		select {
		case <-quit:
			return
		default:
			fmt.Println("0")
		}
	}
}