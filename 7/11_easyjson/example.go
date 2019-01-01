package main

import "fmt"

// Установка
// go get -u github.com/mailru/easyjson/...

//easyjson:json
type JSONData struct {
	Data []string
}

func main() {
	var d JSONData
	input := []byte(`{"Data" : ["One", "Two", "Three"]}`)

	err := d.UnmarshalJSON(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data: %+v", d)
}
