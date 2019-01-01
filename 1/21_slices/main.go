package main

import "fmt"

func main() {
	var sl []int
	fmt.Println("пустой слайс", sl)
	sl = append(sl, 100) // добавление элемента в слайс
	fmt.Println("уже не пустой слайс", sl)

	fmt.Println("длина слайса", len(sl))

	fmt.Println("длина внутреннего массива в слайсе", sl, cap(sl))
	sl = append(sl, 102)
	fmt.Println("длина внутреннего массива в слайсе", sl, cap(sl))
	sl = append(sl, 103)
	sl = append(sl, 104)
	fmt.Println("длина внутреннего массива в слайсе", sl, cap(sl))
	sl = append(sl, 105)
	fmt.Println("длина внутреннего массива в слайсе", sl, cap(sl))

	// короткая инициализация
	sl2 := []int{10, 20, 30}
	fmt.Println(sl2)

	// добавить слайс в слайс
	sl = append(sl, sl2...)
	fmt.Println(sl)

	// создать слайс с нужной длиной сразу
	slice3 := make([]int, 10)
	fmt.Println(slice3, len(slice3), cap(slice3))

	// создать слайс с нужной длиной и размером
	slice4 := make([]int, 10, 15)
	fmt.Println(slice4, len(slice4), cap(slice4))
	slice4 = append(slice4, []int{1, 2, 3, 4, 5, 6}...)
	fmt.Println(slice4, len(slice4), cap(slice4))

	// внутри слайса - ссылка на массив, она копируется если просто присвоить
	slice5 := slice4
	slice5[1] = 100500
	fmt.Println(slice4, slice5)

	//неправильная попытка скопировать слайс - он оппробует скопировать сколько влезет
	var slice6 []int
	copy(slice6, slice5)
	fmt.Println(slice6)

	//правильное копирование слайса
	slice7 := make([]int, len(slice5), len(slice5))
	copy(slice7, slice5)
	fmt.Println(slice7)

	fmt.Println("часть слайса", slice7[1:5], slice7[:2], slice7[10:])
	slice8 := append(slice7[:2], slice7[10:]...)
	fmt.Println("из кусков слайса", slice8)

	a := [...]int{5, 6, 7}
	sl8 := a[:]
	a[1] = 8
	fmt.Println("слайс из массива", sl8)
}
