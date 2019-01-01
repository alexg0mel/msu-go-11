package handlers

import (
	"net/http"
	"io/ioutil"
	"log"
)

func (h Handler) HandleIndex(w http.ResponseWriter, r *http.Request) {
	fileContents, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Write(fileContents)
}
