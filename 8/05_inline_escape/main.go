package main

import "fmt"

const width, height = 640, 480

type Cursor struct {
	X, Y int
}

func Center(c *Cursor) {
	c.X = width / 2
	c.Y = height / 2
}

func CenterCursor() {
	c := new(Cursor)
	Center(c)

	fmt.Println(c.X, c.Y)
}

func main()  {
  // go build -gcflags=-m main.go
}