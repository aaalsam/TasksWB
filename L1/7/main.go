package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map

type Storage struct { // Представляем новую структуру, которая предоставляет собственный мьютек
	sync.RWMutex // Взаимное исключени, защита доступа к памяти
	map1 map[int]int
}

func main() {

	val := &Storage{ // Создаем объект структуры
		map1: map[int]int{},
	}

	writeAndReadMap(val)

}

func writeAndReadMap(val *Storage) {

	wg := new(sync.WaitGroup) // Создаем экземпляр WaitGroup, для того чтобы дождаться завершения горутин в один момент времени

	for i := 0; i < 5; i++ {
		wg.Add(1)        // Добавляем в группу одну горутину
		go func(i int) { // Вызываем анонимную функцию в горутине
			defer wg.Done() // После выполнения блока кода функции, вызываем Done, при котором счётчик уменьшается на единицу
			val.Set(i)      // Добавляем значение i в мапу
		}(i)
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			num := val.Get(i) 
			fmt.Println(num)
		}(i)
	}

	wg.Wait() // Условие, ждущие, пока переменная wg станет нулевой и затем выходит из цикла
}

func (st *Storage) Set(num int) { // Функция добавления значений в мапу
	st.Lock()          // Блокировка на запись
	defer st.Unlock()  // Разблокирока на запись после выполнения блока кода функции
	st.map1[num] = num // Запись по ключу num значения num
}

func (st *Storage) Get(num int) int {
	st.RLock()         // Блокировка на чтение
	defer st.RUnlock() // Разблокирока на чтение после выполнения блока кода функции
	number := st.map1[num]
	return number
}
