// Но гораздо проще объединять логику каналов через оператор select
// Он позволяет проверить на возможность операции с несколькими каналами сразу.
// Важно, если нет доступных вариантов и нет блока default, подпрограмма заблокируется
// Если доступны для работы более одного канала, выбирается произвольный
package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем пару каналов
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			// time.After возвращает канал, в который запись произойдет через 1 секунду
			case <-time.After(time.Second):
				fmt.Println("timeout")
			}
		}
	}()

	fmt.Scanln()
}
