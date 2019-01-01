// Иногда у вас есть несколько каналов, и вы хотите читать из любого из них
// В самом простом случае - можно сделать функцию, которая вернет значения из любого переданного канала
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
func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
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
