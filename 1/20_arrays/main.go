package main

import "fmt"

const size uint = 3

func main() {
	// инициализируется значениями по-умолчанию
	var a1 [3]int
	fmt.Println("массив", a1, "длина", len(a1))

	// можно использовать типизированную беззнаковую константу
	var a2 [2 * size]bool
	fmt.Println(a2, "длина", len(a2))

	a3 := [...]int{1, 2, 3}
	fmt.Println("длина при инициализации", a3, "длина", len(a3))

	a3[1] = 12
	fmt.Println("после изменения", a3)

	// нельзя, проверка при компиляции
	// a3[4] = 12
	// invalid array index 4 (out of bounds for 3-element array)

	var aa [3][3]int
	aa[1][1] = 1
	fmt.Println("массив массивов", aa)
}
