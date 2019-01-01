package main

import (
	"fmt"
	"bytes"
	"sync"
)

var pool map[int]*sync.Pool

func PowerOfTwo(v int) int {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++

	return v
}

func JsToJson(js []byte) []byte {
	colons := bytes.Count(js, []byte{':'})
	resLen := colons*2 + len(js)
	json := pool[PowerOfTwo(resLen)].Get().([]byte)
	defer pool[cap(json)].Put(json[:0])

	const (
		COLON       = byte(':')
		OPEN_BRACE  = byte('{')
		CLOSE_BRACE = byte('}')
		COMMA       = byte(',')
		QUOTE       = byte('"')
		SPACE       = byte(' ')
		KEY_START   = iota
		KEY_END
		VALUE_START
		VALUE_END
		NONE
	)

	CURRENT := NONE

	for i := 0; i < len(js); i++ {
		switch js[i] {
		case OPEN_BRACE:
			CURRENT = KEY_START
			json = append(json, js[i])
			if js[i+1] == SPACE {
				continue
			}
			CURRENT = KEY_START
			json = append(json, QUOTE)
		case COLON:
			if CURRENT == KEY_START {
				json = append(json, QUOTE)
				json = append(json, js[i])
			}
			CURRENT = KEY_END
		case COMMA:
			CURRENT = KEY_START
			json = append(json, js[i])
			if js[i+1] == SPACE {
				continue
			}
		case SPACE:
			if js[i+1] == SPACE {
				continue
			}

			if CURRENT == KEY_START {
				json = append(json, QUOTE)
			}

		default:
			json = append(json, js[i])
		}
	}

	return json
}

func Init() {
	var minSize = 256
	var maxSize = 65536
	pool = make(map[int]*sync.Pool)
	for i := minSize; i <= maxSize; i = i * 2 {
		func(i int) {
			pool[i] = &sync.Pool{
				New: func() interface{} {
					return make([]byte, 0, i)
				},
			}
		}(i)
	}
}

func main() {
	Init()
	var js = []byte(`{common: {test_key:"0", other_test_key: "somevalue", hello: "6"}}`)

	fmt.Println(string(JsToJson(js)))
}
