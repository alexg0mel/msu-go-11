package main

import "fmt"

func main() {
	var mm map[string]string
	fmt.Println("uninitialized map", mm)
	// panic: assignment to entry in nil map
	// mm["test"] = "ok"

	// полная инициализация
	// var mm2 map[string]string = map[string]string{}
	mm2 := map[string]string{}
	mm2["test"] = "ok"
	fmt.Println(mm2)

	// короткая инициализация
	var mm3 = make(map[string]string)
	mm3["firstName"] = "Vasily"
	fmt.Println(mm3)

	// получение значения
	firstName := mm3["firstName"]
	fmt.Println("firstName", firstName, len(firstName))

	// есть обратиться к несуществующему ключу - отдасться значение по-умолчанию
	lastName := mm3["lastName"]
	fmt.Println("lastName", lastName, len(lastName))

	// проверка на то, что значение есть
	lastName, ok := mm3["lastName"]
	fmt.Println("lastName is", lastName, "exist:", ok)

	// только получение признака существования
	_, exist := mm3["firstName"]
	fmt.Println("fistName exist:", exist)

	// удаление значения
	delete(mm3, "firstName")
	_, exist = mm3["firstName"]
	fmt.Println("fistName exist:", exist)

}
