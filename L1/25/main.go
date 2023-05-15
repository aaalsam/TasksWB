package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func sleep1(sec int) {
	<-time.After(time.Duration(sec) * time.Second) // Оправляет текущее время в канал по истечению установленного времени, затем функция читает из канала и завершается
}

func main() {
	fmt.Println("start")
	sleep1(3)
	fmt.Println("end")
}
