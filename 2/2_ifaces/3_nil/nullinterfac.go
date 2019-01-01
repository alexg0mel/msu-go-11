package main

import "fmt"

func main() {
	var z interface{}

	fmt.Printf("%v %v\n", z, z == nil)

	if f := getNil(10); f != nil {
		fmt.Println("i'm not nill")
	}

	if f := getNil("2nd"); f == nil {
		fmt.Println("but i actually am nil")
	}
}

type myError string

func (m myError) Error() string {
	return string(m)
}

func getNil(input interface{}) error {
	var m *myError
	if _, ok := input.(int); ok {
		return m
	}
	return nil
}

func getTrueNil() error {
	return nil
}
