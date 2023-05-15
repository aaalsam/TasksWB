package main

import (
	"fmt"
)

// Реализовать пересечение двух неупорядоченных множеств

func main() {
	set := make(map[int]struct{})

	set[0] = struct{}{}
	set[5] = struct{}{}
	set[2] = struct{}{}

	set1 := make(map[int]struct{})

	set1[0] = struct{}{} 
	set1[1] = struct{}{}
	set1[2] = struct{}{}

	fmt.Println(intersectionSets(set, set1))

}

func intersectionSets(set map[int]struct{}, set1 map[int]struct{}) *map[int]struct{} { // Функция возвращает пересечение двух множеств

	set2 := make(map[int]struct{})

	if len(set) > len(set1) { // Проверяем по тому сету, который меньше размером
		set, set1 = set1, set
	}

	for e := range set {
		if _, ok := set1[e]; ok { // В set2 добавляется элемент из set, если он содержится в set1
			set2[e] = struct{}{}
		}
	}

	return &set2
}
