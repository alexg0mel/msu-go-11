// Пока в примерах нам было не так важно, что все горутины закончили работу,
// тк при выходе основного процесса, все горутины завершатся,
// но, что если нам требуется удостовериться в завершении всех проивзодимых работ
// Можно это сделать с помощью канала
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(die chan bool) <-chan string { // Возвращаем канал строк только для чтения.
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("boring %d", rand.Intn(100)):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-die:
				fmt.Println("Jobs done!")
				die <- true
				return
			}
		}
	}()
	return c
}

func main() {
	die := make(chan bool)
	res := boring(die)

	for i := 0; i < 5; i++ {
		// Читаем из канала
		fmt.Printf("You say: %q\n", <-res)
	}
	die <- true
	// Ждем, пока все горутины закончат выполняться
	<-die
}
