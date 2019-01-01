package main

import (
	"fmt"
	"net"
	"bufio"
)

func main() {
	// Bind на порт ОС
	listener, _ := net.Listen("tcp", ":5000")

	for {
		// ждём пока не придёт клиент
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("Can not connect!!")
			conn.Close()
			continue
		}

		fmt.Println("Connected")

		// создаём Reader для чтения информации из сокета
		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading")

			//defer conn.Close()

		for {
			// побайтово читаем
			rbyte, err := bufReader.ReadByte()

			if err != nil {
				fmt.Println("Can not read!", err)
				break
			}

			fmt.Print(string(rbyte))
		}
	}
}
