package main

import "fmt"

// example represents a type with different fields.
type example struct {
	Flag    bool
	counter int16
	pi      float32
}

func main() {

	// Declare a variable of type example set to its
	// zero value.
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of type example and init using
	// a struct literal.
	e2 := example{
		Flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	e3 := example{}

	// Display the field values.
	fmt.Println("Flag", e2.Flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)
}
