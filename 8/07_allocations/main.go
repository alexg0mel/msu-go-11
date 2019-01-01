package main

// профалинг бенчмарков
// go test -bench=. -benchmem -memprofile=mem.out
// go tool pprof -alloc_space 07_allocations.test mem.out

import (
	"fmt"
)

func JsToJson(js []byte) []byte {
	json := []byte{}
	//json := make([]byte, len(js))

	//colons := bytes.Count(js, []byte{':'})
	//resLen := colons*2 + len(js)
	//json := make([]byte, 0, resLen)

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

func main() {
	var js = []byte(`{common: {test_key:"0", other_test_key: "somevalue", hello: "6"}}`)

	fmt.Println(string(JsToJson(js)))
}
