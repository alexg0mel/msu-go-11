// Но, безопаснее использовать пакет sync, в частности структуру WaitGroup
// У нее нет публичных полей, но есть 3 метода
// Add увеличивает счетчик ожидаемых работ, Done декрементит,
// Wait - блокируется, пока внутренний счетчик не станет равным 0
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func boring1(wg *sync.WaitGroup, die chan bool) <-chan string { // Возвращаем канал строк только для чтения.
	c := make(chan string)
	go func() {
		for {
			select {
			case c <- fmt.Sprintf("boring %d", rand.Intn(100)):
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			case <-die:
				fmt.Println("Jobs done!")
				wg.Done()
				return
			}
		}
	}()
	return c
}

func main() {
	die := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	res1 := boring1(&wg, die)
	res2 := boring1(&wg, die)

	for i := 0; i < 5; i++ {
		// Читаем из канала
		fmt.Printf("1st say: %q\n", <-res1)
		fmt.Printf("2nd say: %q\n", <-res2)
	}
	die <- true
	// Ждем, пока все горутины закончат выполняться
	wg.Wait()
}
