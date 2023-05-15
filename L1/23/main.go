package main

import "fmt"

// Удалить i-ый элемент из слайса

func main() {

	sl := make([]int, 5)
	sl[0] = 1
	sl[1] = 2
	sl[2] = 3
	sl[3] = 4
	sl[4] = 5

	fmt.Println(sl)

	i := 2

	removeI(sl, i)

	remove(sl, i)

	fmt.Println(sl)
}

func removeI(sl []int, i int) {
	sl[i] = sl[len(sl)-1] // копируем последний элемент в индекс i
	sl = sl[:len(sl)-1]   // Урезаем слайс
	fmt.Println(sl)
}

func remove (slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}
