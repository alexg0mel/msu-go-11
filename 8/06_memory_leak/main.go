package main

// профилируем память
// go tool pprof -alloc_space main http://127.0.0.1:8080/debug/pprof/heap\?seconds\=20
// main - путь к бинарнику

import (
	"net/http"
	_ "net/http/pprof" // пвстраивает проилировщик в программу
)

func WaitingFunc(c chan struct{}) {
	<-c
}

func Handler(w http.ResponseWriter, r *http.Request) {
	c := make(chan struct{}, 3)
	go WaitingFunc(c)

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", Handler)

	http.ListenAndServe(":8080", nil)
}
