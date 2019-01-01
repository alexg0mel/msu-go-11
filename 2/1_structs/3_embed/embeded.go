package main

import "fmt"

type Person struct {
	Name string
	inn  string
}

type Stuff struct {
	inn int
}

type SecretAgent struct {
	Person
	Stuff
	LicenseToKill bool
}

func (p Person) GetName() string {
	return p.Name
}

func (s SecretAgent) GetName() string {
	return "CLASSIFIED"
}

func main() {
	sa := SecretAgent{Person: Person{"James", "12312321321"}, LicenseToKill: true}

	//fmt.Printf("%T %+v\n", sa, sa)
	fmt.Println("secret inn", sa.GetName())
}
