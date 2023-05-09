package main

import (
	"fmt"
	"reflect"
)

// Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func main() {

	var i interface{} = make(chan int)

	fmt.Println(reflect.TypeOf(i))

	var i1 interface {} = 2

	definitionType(i1)

}

func definitionType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown type")
	}
}
