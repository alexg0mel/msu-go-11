package main

import (
	"errors"
	"fmt"
)

// Create a named type for our new error type.
type errorString string

// Implement the error interface.
func (e errorString) Error() string {
	return string(e)
}

// New creates interface values of type error.
func New(text string) error {
	return errorString(text)
}

var ErrNamedType = New("EOF")
var ErrStructType = errors.New("EOF")

func main() {
	if ErrNamedType == New("EOF") {
		fmt.Println("Named Type Error")
	}

	if ErrStructType == errors.New("EOF") {
		fmt.Println("Struct Type Error")
	}

	err := BadFunc()

	switch err := err.(type) {
	case nil:
		// call succeeded, nothing to do
	case *MyError:
		fmt.Println("error occurred on line:", err.Line)
	default:
		// unknown error
	}

	st := []int{10, 12, 15}
	ist := make([]interface{}, len(st))
	for i := range st {
		ist[i] = st[i]
	}

	DoIfaces(ist...)
}

type MyError struct {
	Msg  string
	File string
	Line int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.File, e.Line, e.Msg)
}

func BadFunc() error {
	return &MyError{"Something happened", "server.go", 42}
}

type temporary interface {
	Temporary() bool
}

// IsTemporary returns true if err is temporary.
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

func DoIfaces(slice ...interface{}) error {
	return nil
}
