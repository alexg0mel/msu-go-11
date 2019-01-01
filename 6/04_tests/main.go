package main

import "github.com/pkg/errors"

func main() {

}

func Sum(a int, b int) int {
	return a + b
}

func Division(a int, b int) int {
	return a / b
}

func TrueDivision(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}

	return a / b, nil
}

type User struct {
	Name string
	Age int
}

func AddUser(users *[]User, name string, age int) {
	user := User{
		Name: name,
		Age: age,
	}

	*users = append(*users, user)
}

