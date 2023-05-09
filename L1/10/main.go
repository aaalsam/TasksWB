package main

import (
	"fmt"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов.
// Последовательность в подмножноствах не важна. Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func TmGroup(tmp []float64) map[int][]float64 {

	groups := make(map[int][]float64) // в ключах будут храниться группы с шагом в 10, а в значениях слайсы чисел с плавающей точкой

	for _, val := range tmp { // Итерация по слайсу температур
		func(val float64) {
			if val > 0 {
				for i := float64(0); i < 100; i += 10 {
					if i <= val && val < i+10 {
						groups[int(i)] = append(groups[int(i)], val)
						break
					}
				}
			} else {
				for i := float64(0); i > -100; i -= 10 {
					if i >= val && val > i-10 {
						groups[int(i)] = append(groups[int(i)], val)
						break
					}
				}
			}
		}(val)
	}
	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := TmGroup(temperatures)

	fmt.Println(groups)
}
