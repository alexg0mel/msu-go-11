package main

const someInt = 1
const typedInt int32 = 17
const fullName = "Vasily"

const (
	flagKey1 = 1
	flagKey2 = 2
)

const (
	one = iota
	two
	_    // пустая переменная, пропуск iota
	four // = 4
)

const (
	_         = iota // пропускаем первое значне
	KB uint64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	// ошибка! переполнение типа
	// ZB
)

func main() {
	pi := 3.14

	// тип константы может быть определён во время компиляции
	println(pi + someInt)

	// константа может иметь тип
	// println(pi + typedInt)
	// invalid operation: pi + typedInt (mismatched types float64 and int32)

	println(KB, MB, GB, TB, PB, EB)
}
