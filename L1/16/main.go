package main

import (
	"fmt"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.


func quickSort(a []int, left int, right int) []int{
	if left < right {
		pivot := a[left]
		//fmt.Println("pivot", pivot)

		i, j := left, right // Заводим два счётчика: left идёт слева, right идёт справа
		for {
			for a[j] >= pivot && i < j { // Уменьшаем j пока элементы больше пивота
				//fmt.Println("rigth", a[j])
				j--
			}
 
			for a[i] <= pivot && i < j {  // Увеличиваем i пока элементы меньше пивота
				//fmt.Println("left", a[i])
				i++
			}

			if i >= j { // когда счетчики встретятся выходим из цикла
				break
			}
 
			a[i], a[j] = a[j], a[i] // меняем местами элементы
		}

		a[left] = a[i]
		a[i] = pivot
 
		quickSort(a, left, i-1) // Рекурсивно сортируется отдельно слева и правые стороны
		quickSort(a, i+1, right)
	}
	return a
}
 
func main() {
	a := []int{3, 7, 2, 9, 0, 1, 6, 8, 4} // 8
	fmt.Println(a)
	fmt.Println(quickSort(a, 0, len(a) - 1))
}
