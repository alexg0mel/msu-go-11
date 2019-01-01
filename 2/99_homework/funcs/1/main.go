package main

import "fmt"

type memoizeFunction func(int, ...int) interface{}

// TODO реализовать
var fibonacci memoizeFunction
var romanForDecimal memoizeFunction

//TODO Write memoization function

func memoize(function memoizeFunction) memoizeFunction {
	cache := make(map[int] interface{})
	return func(x int, xn ...int) interface{} {
		if value, found := cache[x]; found {
			return value
		}
		value := function(x, xn...)
		cache[x] = value
		return value
	}
}

// TODO обернуть функции fibonacci и roman в memoize
func init() {
	fibonacci = memoize(func(x int, xn ...int) interface{} {
		if x < 2 {return x}
		return fibonacci(x-1).(int) + fibonacci(x-2).(int)
	})
	decimals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X",
		"IX", "V", "IV", "I"}


	romanForDecimal = memoize(func(x int, xn ...int) interface{} {
		var res string;
		var i int
		nextX := x
		if nextX <=0 {return res}
		for  _, key := range decimals {
			if nextX == key {
				nextX -=key
				res+=romans[i]

				break
			}
			if nextX - key >0 {
				nextX -=key
				res += romanForDecimal(key).(string)
				break
			}
			i++
		}

		return res + romanForDecimal(nextX).(string)

	})

}

func main() {
	fmt.Println("Fibonacci(45) =", fibonacci(45).(int))
	for _, x := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
		14, 15, 16, 17, 18, 19, 20, 25, 30, 40, 50, 60, 69, 70, 80,
		90, 99, 100, 200, 300, 400, 500, 600, 666, 700, 800, 900,
		1000, 1009, 1444, 1666, 1945, 1997, 1999, 2000, 2008, 2010,
		2012, 2500, 3000, 3999} {
		fmt.Printf("%4d = %s\n", x, romanForDecimal(x).(string))
	}
}
