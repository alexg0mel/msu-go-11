package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func Readfile(f string) ([]byte, error) {
	file, err := os.OpenFile(f, 0, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return nil, nil
}
