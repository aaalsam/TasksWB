package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter1 struct {
	num uint64
}

func (c *Counter1) Inc() { //метод структуры счетчика, прибавляющая значение к счетчику

	atomic.AddUint64(&c.num, 1) // Для того, чтобы доступ к счетчику был безопасен, чтобы только одна горутина смогла выполнять действие со счетчиком

}

func (c *Counter1) Value() uint64 { // метод возвращает значение счетчика
	return c.num
}

func main() {

	wg := new(sync.WaitGroup)

	cnt := &Counter1{ //Объявляем и инициализируем счетчик
		num: 0,
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt.Inc()
		}()
	}

	wg.Wait()

	fmt.Println(cnt.Value())
}