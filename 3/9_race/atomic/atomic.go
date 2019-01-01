// Также, когда само общее значение - это просто число
// Можно использовать пакет atomic, для того, чтобы гарантировать очередность изменений объекта
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type atomicCounter struct {
	val int64
}

func (c *atomicCounter) Add(x int64) {
	atomic.AddInt64(&c.val, x)

}

func (c *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	counter := atomicCounter{}

	// Если запустить этот код с ключом race, можно заметить, что мы никак не гарантируем
	// Завершение всех работ, для того, чтобы гарантии были, стоит использовать WaitGroup
	// var wg sync.WaitGroup
	// wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(no int) {
			// Важно отметить, что в рамках этого цикла, управление между
			// горутинами не может переключаться, если мы хотим, чтобы работа шла более паралельно
			// потребуется использовать вызов Gosched (управление переключается только на операторе select и работе с ОС, такой как чтение из файлов или сети)
			for i := 0; i < 10000; i++ {
				counter.Add(1)
				//runtime.Gosched()
			}
			//wg.Done()
		}(i)
	}

	time.Sleep(time.Second)
	// wg.Wait()
	fmt.Println(counter.Value())
}
