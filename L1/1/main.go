package main

import "fmt"

// Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

// Объявляем именованный тип
type Human struct {
	Name string
}

// Создаем метод типа Human
func (h *Human) printName() {
	fmt.Println(h.Name)
}

// Объявдяем именованный тип и встраиваем в его поле тип Human
type Action struct {
	Human
}

func main() {
	ira := Human{Name: "Ira"}
	
	action := Action{ira} // Передаем в Action экземпляр структуры Human

	action.Human.printName() // полный путь
	action.printName()       // shorthand. Go проверяет наличие метода printName, начиная от корня типа Action углубляясь дальше - во вложенный типы, и берет ближайшее.
}