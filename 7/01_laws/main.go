package main

// reflect.Type reflect.Value

import (
	"fmt"
	"reflect"
)

func main() {
	x := 3.4
	// TypeOf() принимает на вход interface{}, в этом месте будет аллокация
	fmt.Println("reflect.Type:", reflect.TypeOf(x))

	// reflect.Value != значению переданному на вход
	fmt.Println("reflect.Value:", reflect.ValueOf(x).String())

	v := reflect.ValueOf(x)
	fmt.Println("Тип value:", v.Type())
	fmt.Println("тип float64:", v.Kind() == reflect.Float64)
	fmt.Println("Значение:", v.Float())

	// Важно отметить, что Kind - это базовый тип (int/float/struct/slice)
	// А не пользовательский тип
	type MyInt int
	var c MyInt = 7
	v = reflect.ValueOf(c)
	fmt.Println("kind is int: ", v.Kind() == reflect.Int) // true.

	y := v.Interface().(float64) // y will have type float64.
	fmt.Println("Значение обертки", v, "Само значение", y)

	// access()
}

func access() {
	var x float64 = 3.4
	// Мы создаем копию
	v := reflect.ValueOf(x)
	// Изменeние v запрещено, тк отсутствует связь с подлежащим значением
	v.SetFloat(7.1) // Error: will panic.

	fmt.Println("settability of v:", v.CanSet())

	// Чтобы иметь возможность изменить значение, нам потребуется ссылка
	p := reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of p:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	// Теперь, использую Elem мы получим Value, лежащее по ссылке
	v = p.Elem()
	fmt.Println("settability of v:", v.CanSet())

	v.SetFloat(7.1)
	fmt.Println(v.Interface())
	fmt.Println(x)
}
