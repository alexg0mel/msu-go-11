package main

import (
	"encoding/json"
	"fmt"
	"msu-go-11/8/04_json/structs"
)

func main() {
	data := structs.Json{
		Tasks: []structs.Task{
			{Title: "first"},
			{Title: "second"},
			{Title: "third"},
		},
		Auth: struct {
			SessionID string `json:"session_id"`
			CsrfToken string `json:"csrf_token"`
		}{
			SessionID: "423j345k34h5lh425h34k54343454353rfcwedf",
			CsrfToken: "jflkwnlhtl24hrt3kl4t3k4bt34",
		},
	}

	str := string(Reflect(data))
	fmt.Println(str)
}

func Reflect(data structs.Json) []byte {
	bytes, _ := json.Marshal(data)
	return bytes
}

func Generated(data structs.Json) []byte {
	bytes, _ := data.MarshalJSON()
	return bytes
}
