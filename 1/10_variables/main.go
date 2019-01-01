/*многострочный комментарий
переменные в golang */
package main

func main() {
	// целые числа
	var i int = 10 // платформозависимый тип, 32 или 64 бита
	var autoInt = -10
	var bigInt int64 = 1<<32 - 1          // int8, int16, int32, int64
	var unsignedInt uint = 100500         // платформозависимый тип, 32 или 64 бита
	var unsignedBigInt uint64 = 1<<64 - 1 // uint8, unit16, uint32, unit64
	println("integers", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// числа с плавающей точкой
	var p float32 = 3.14 // float = float32, float64
	println("float: ", p)

	//булевые переменные
	var b = true
	println("bool variable", b)

	//строки
	var hello string = "Hello\n\t"
	var world = "World"
	println(hello, world)

	//бинарные данные
	var rawBinary byte = '\x27'
	println("rawBinary", rawBinary)

	// так нельзя
	// var singleQuote string = 'Hello world'
	// missing '
	// syntax error: unexpected ello at end of statement

	/*
		короткое объявление
	*/
	meaningOfLive := 42
	println("Meaning of life is ", meaningOfLive)
	// работает только для новых переменных, world объявлен выше, поэтому ошибка
	// world := "Мир"
	// no new variables on left side of :=

	/*
		приведение типов
	*/
	println("float to int conversion ", int(p))
	println("int to string conversion ", string(48))

	// комплексные числа
	z := 2 + 3i
	println("complex number: ", z)

	/*
		операции со строками
	*/
	s1 := "Vasily"
	s2 := "Romanov"
	fullName := s1 + s2
	println("name length is: ", fullName, len(fullName))

	escaping := `Hello\r\n
	World`
	println("as-is escaping: ", escaping)

	/*
		значение по-умолчанию
	*/
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println("default values: ", defaultInt, defaultFloat, defaultString, defaultBool)

	/*
		несколько переменных
	*/
	var v1, v2 string = "v1", "v2"
	println(v1, v2)

	var (
		m0 int = 12
		m2     = "string"
		m3     = 23
	)
	println(m0, m2, m3)

}
