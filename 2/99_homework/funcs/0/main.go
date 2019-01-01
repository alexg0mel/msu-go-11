package main

import (
	"fmt"
)

// TODO: Реализовать вычисление Квадратного корня
func Sqrt(x float64) float64 {
	res := x
	// Итерация Герона
	for i:=0;i<4;i++ {
		res = 0.5 * (res + x / res)
	}
	return res
}

func main() {
	fmt.Println(Sqrt(2))
}
