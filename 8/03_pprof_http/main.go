package main

import (
	"net/http"
	_ "net/http/pprof"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

// go tool pprof main http://127.0.0.1:8080/debug/pprof/profile\?seconds\=5

// cd $GOPATH/src/github.com/uber/go-torch
// git clone https://github.com/brendangregg/FlameGraph.git
