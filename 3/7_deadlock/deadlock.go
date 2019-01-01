// Если при запуске программы все go рутины окажутся в состоянии ожидания,
// runtime выкинет панику, с сообщением all goroutines are sleeping
package main

import (
	"fmt"
	"time"
)

// Ball is just a ball
type Ball struct{ hits int }

func main() {
	// Создаем канал для взаимодействия игроков
	table := make(chan *Ball)
	// Старутем пару игроков
	go player("ping", table)
	go player("pong", table)

	// table <- new(Ball) // Запуска мяча в игру
	time.Sleep(1 * time.Second)
	<-table // Конец игры, забираем мяч
}

func player(name string, table chan *Ball) {
	for {
		// Ждем, когда мяч попал к игроку
		ball := <-table
		// Увеличиваем счетчик ударов
		ball.hits++
		fmt.Println(name, ball.hits)
		// Ждем немного
		time.Sleep(100 * time.Millisecond)
		// Отправляем мяч обратно в канал
		// Важно, программа заблокируется, пока другой игрок оттуда не прочитает
		table <- ball
	}
}
