// Гораздо удобнее объединять каналы через select

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring; I'm leaving.")
}

// Принимаем 2 канала строк, из которых можем только читать
// возвращаем 1 канал только для чтения
// Теперь достаточно только одной горутины
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func boring(msg string) <-chan string { // Возвращаем канал строк только для чтения.
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		}
	}()
	return c
}
