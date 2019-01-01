package pipeline

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	res := []interface{}{}
	case1 := []job{
		job(func(in, out chan interface{}) {
			for i := 0; i < 10; i++ {
				out <- i
			}
		}),
		job(func(in, out chan interface{}) {
			for i := range in {
				if s, ok := i.(int); ok && (s%3 == 0) {
					out <- i
				}
			}
		}), job(func(in, out chan interface{}) {
			for value := range in {
				out <- interface{}(value.(int) * 10)
			}
		}), job(func(in, out chan interface{}) {
			for val := range in {
				res = append(res, val)
			}
		}),
	}

	expected := []interface{}{0, 30, 60, 90}

	Pipe(case1...)
	for r := range res {
		if !reflect.DeepEqual(res[r], expected[r]) {
			t.Errorf("Failed output")
			t.FailNow()
		}
	}

	res2 := ""

	case2 := []job{
		job(func(in, out chan interface{}) {
			for _, word := range []string{"Hello", "World"} {
				out <- word
			}
		}),
		job(func(in, out chan interface{}) {
			for word := range in {
				fmt.Println("Got", word)
				if w, ok := word.(string); ok {
					fmt.Println("And it is a string")
					if len(res2) > 0 {
						res2 = res2 + " " + w
					} else {
						res2 = w
					}
				}
			}
		}),
	}

	Pipe(case2...)
	expected2 := "Hello World"

	if !reflect.DeepEqual(res2, expected2) {
		t.Errorf("Failed output, expected %s got %s", expected2, res2)
		t.FailNow()
	}

}
