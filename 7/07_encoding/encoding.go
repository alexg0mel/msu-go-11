package main

// Примеры работы encoding json

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type App struct {
	Id    string `json:"id"`
	Title string `json:"title,omitempty"`
	// Все не экспортируемые поля не учавствуют в маршаллинге
	hidden string
}

func main() {

	data := []byte(`
    {
        "id": "k34rAT4",
        "title": "My Awesome App"
    }
`)

	var app App
	err := json.Unmarshal(data, &app)

	if err != nil {
		panic(err)
	}

	fmt.Printf("App: %+v\n", app)

	// embedding()
	// pointers()
	// custom()
}

// Правила вложенности сохраняются

type App1 struct {
	Id string `json:"id"`
}

type Org struct {
	Name string `json:"name"`
}

type AppWithOrg struct {
	App1
	Org
}

func embedding() {
	data := []byte(`
    {
        "id": "k34rAT4",
        "name": "My Awesome Org"
    }
`)

	var appWithOrg AppWithOrg
	err := json.Unmarshal(data, &appWithOrg)
	if err != nil {
		panic(err)
	}

	app := appWithOrg.App1
	org := appWithOrg.Org

	fmt.Printf("App1: %+v,\n Org: %+v\n", app, org)
}

// Можно не использовать структуры, а разбираться уже потом,
// Но, опять же, мы теряем все compile time проверки
// и придётся далее угадывать подлежащие структуры

// https://golang.org/pkg/encoding/json/#RawMessage

func pointers() {
	var parsed map[string]interface{}

	data := []byte(`
    {
        "id": "k34rAT4",
        "age": 24
    }
`)

	err := json.Unmarshal(data, &parsed)
	if err != nil {
		panic(err)
	}
	fmt.Println("Id:", parsed["id"].(string))
}

/*
Cheatsheet по преобразованию
bool, for JSON booleans
float64, for JSON numbers
string, for JSON strings
[]interface{}, for JSON arrays
map[string]interface{}, for JSON objects
nil for JSON null
*/

// Пример кастомного маршаллера
func custom() {
	data := []byte(`{ "Month": "04/2018" }`)
	var d day

	err := json.Unmarshal(data, &d)
	if err != nil {
		panic(err)
	}
	fmt.Println("Got month:", d.Month)
}

type day struct {
	Month Month
}

type Month struct {
	MonthNumber int64
	YearNumber  int64
}

func (m Month) String() string {
	return fmt.Sprintf("%d/%d", m.MonthNumber, m.YearNumber)
}

func (m Month) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m *Month) UnmarshalJSON(value []byte) (err error) {
	if len(value) < 2 {
		return fmt.Errorf("bad data")
	}

	// Чистим кавычки
	if value[0] == '"' {
		value = append(value[:0], value[1:]...)
	}

	if value[len(value)-1] == '"' {
		value = value[0 : len(value)-1]
	}

	parts := strings.Split(string(value), "/")

	m.MonthNumber, err = strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return
	}
	m.YearNumber, err = strconv.ParseInt(parts[1], 10, 32)
	if err != nil {
		return
	}

	return nil
}
