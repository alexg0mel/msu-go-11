package main

import "fmt"

func main() {
	showMeTheMoney()

	stuff := []int{10, 2, 3}
	res := sumMore(stuff...)

	fmt.Println("Res", res)
}

func showMeTheMoney() {
	fmt.Printf("$$$$")

}

// Несколько входящих значений
func sum(i int, j int) int {
	return i + j
}

// Упрощенная запись для нескольких значений одинакового типа
func sumLight(i, j int) int {
	return i + j
}

// Для получения произвольного списка однотипных значений
func sumMore(stuff ...int) (res int) {
	for i := range stuff {
		res += stuff[i]
	}
	return
}

// Возврат нескольких значений
func sumOnlyNatural(stuff ...int) (int, error) {
	res := 0
	for i := range stuff {
		if stuff[i] < 0 {
			return 0, fmt.Errorf("Only natural numbers expected - given %d", stuff[i])
		}
		res += stuff[i]
	}
	return res, nil
}

// Возврат Именованных значений
func sumNatural2(stuff ...int) (res int, err error) {
	for i := range stuff {
		if stuff[i] < 0 {
			err = fmt.Errorf("Only natural numbers expected - given %d", stuff[i])
			return
		}
		res += stuff[i]
	}
	return res, err
}
