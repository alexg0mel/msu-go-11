package main

import "fmt"

type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

func main() {
	d := Pill(1)

	v := Pill(3)

	fmt.Println(d, v)
}

//go:generate stringer -type=Pill
