package main

// Пакет reflect позволяет творить довольно сумашедшие вещи
// допустим, мы хотим организовать работы с абстактными каналами данных
// Но у этого подхода есть свои проблемы
// Отсутствие compile time проверок типов, одна из важнейших

import (
	"log"
	"reflect"
	"sync"
)

// Инициализация абстрактного канала
func makeChan(chp interface{}, buffsize int) {

	chpv := reflect.ValueOf(chp)
	if chpv.Kind() != reflect.Ptr {
		log.Panic("Первый аргумент должен быть ссылкой на канал")
	}

	chv := chpv.Elem()
	if chv.Kind() != reflect.Chan {
		log.Panic("Первый аргумент должен быть ссылкой на канал")
	}

	chantype := chv.Type()
	chv.Set(reflect.MakeChan(chantype, buffsize))
}

// отправка в в канал произвольного типа
func send(ch interface{}, value interface{}) {
	chv := reflect.ValueOf(ch)
	if chv.Kind() != reflect.Chan {
		log.Panic("first argument must be chan")
	}

	v := reflect.ValueOf(value)
	if chv.Type().Elem() != v.Type() {
		log.Panic("chan and value don't much type")
	}

	if chv.Type().ChanDir() == reflect.RecvDir {
		log.Panic("first argument must be sendable chan")
	}

	chv.Send(v)
}

// Получение данных из произвольного канала
func recv(ch interface{}, p interface{}) bool {
	chv := reflect.ValueOf(ch)
	if chv.Kind() != reflect.Chan {
		log.Panic("first argument must be chan")
	}

	pv := reflect.ValueOf(p)
	if pv.Kind() != reflect.Ptr {
		log.Panic("second argument must be pointer")
	}

	if chv.Type().Elem() != pv.Type().Elem() {
		log.Panic("chan and value don't much type")
	}

	if chv.Type().ChanDir() == reflect.SendDir {
		log.Panic("first argument must be receivable chan")
	}

	v, ok := chv.Recv()
	pv.Elem().Set(v)

	return ok
}

func selectCase(recvCh interface{}, recvCase func(v interface{}, ok bool), sendCh interface{}, sendValue interface{}, sendCase func(), defaultCase func()) {
	recvChv := reflect.ValueOf(recvCh)
	if recvChv.Kind() != reflect.Chan || recvChv.Type().ChanDir() == reflect.SendDir {
		log.Panic("first argument must be receivable chan")
	}

	sendChv := reflect.ValueOf(sendCh)
	if sendChv.Kind() != reflect.Chan || sendChv.Type().ChanDir() == reflect.RecvDir {
		log.Panic("third argument must be sendable chan")
	}

	v := reflect.ValueOf(sendValue)
	if sendChv.Type().Elem() != v.Type() {
		log.Panic("sendCh and sendValue don't much type")
	}

	if sendChv.Type().ChanDir() == reflect.RecvDir {
		log.Panic("first argument must be sendable chan")
	}

	cases := []reflect.SelectCase{
		reflect.SelectCase{
			reflect.SelectRecv,
			recvChv,
			reflect.ValueOf(nil),
		},
		reflect.SelectCase{
			reflect.SelectSend,
			sendChv,
			v,
		},
		reflect.SelectCase{
			reflect.SelectDefault,
			reflect.ValueOf(nil),
			reflect.ValueOf(nil),
		},
	}

	chosen, recv, recvOK := reflect.Select(cases)
	switch chosen {
	case 0:
		recvCase(recv.Interface(), recvOK)
	case 1:
		sendCase()
	case 2:
		defaultCase()
	}
}

func main() {
	var ch chan int
	makeChan(&ch, 10)
	send(ch, 1)
	log.Printf("%d/%d", len(ch), cap(ch))
	send(ch, 2)
	log.Printf("%d/%d", len(ch), cap(ch))
	var n1 int
	recv(ch, &n1)
	log.Println(n1)

	ch2 := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go selectCase(
			// case receive
			ch, func(v interface{}, ok bool) {
				log.Println(v, ok)
				wg.Done()
			},
			// case send
			ch2, 10, func() {
				log.Println("send")
				wg.Done()
			},
			// default
			func() {
				log.Println("default", <-ch2)
				wg.Done()
			},
		)
	}
	wg.Wait()
}
