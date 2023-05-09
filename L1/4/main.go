package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {

	var n int // Создали переменную для сканера

	fmt.Println("Введите N воркеров:")
	fmt.Scan(&n) //Считали данные с консоли, записали в n

	ch := make(chan int)

	osSignal := make(chan os.Signal, 1) // Сигнал SIGINT отправляется при введении пользователем в управляющем терминале символа прерывания, по умолчанию это ^C (Control-C).
	// перехватыват сигнала
	signal.Notify(osSignal, syscall.SIGINT) // Если в osSignal поступит сигнал SIGINT, то в канал osSignal запишется значение

	read(ch, n)

	write(ch, osSignal)
}

func read(ch chan int, n int) {
	for i := 0; i < n; i++ {
		go func() {
			for num := range ch { // Читает из канала, пока он не закрыт
				fmt.Println(num)
			}
		}()
	}
}

func write(ch chan int, osSignal chan os.Signal) {
	for {
		select { // Оператор select позволяет go-процедуре находиться в ожидании нескольких операций передачи данных.
		case <-osSignal:
			close(ch)
			fmt.Println("Завершено")
			return // Выходим из функции
		default:
			ch <- 2
		}
	}
}
