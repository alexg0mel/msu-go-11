package main

import (
	"reflect"
	"strconv"
	"testing"
)

func ReturnInt() int  {
	return 1
}

func ReturnFloat() float32  {
	return 1.1
}

func ReturnIntArray() [3]int  {
	return [3]int {1,3,4}
}

func ReturnIntSlice() []int  {
	return []int {1,2,3}
}

func IntSliceToString(in []int) string  {
	var res string
	for _, el:=range in{
		res+=strconv.Itoa(el)
	}
	return  res
}

func MergeSlices(in_float []float32, in_int[] int32) []int {
	var res []int
	for _, el1 := range in_float{
		res = append(res, int(el1))
	}
	for _, el2:= range in_int{
		res = append(res, int(el2))
	}
	return  res
}

func GetMapValuesSortedByKey(inp map[int]string) []string {
	var res [] string
	var max,min,i int

	for i= range inp{
		if i > max { max = i}
		if min==0 || i < min  {min = i}
	}

	for i=min; i<=max; i++{
		e,q:= inp[i]
		if(q) {res = append(res,e)}
	}
	return res
}

func TestReturnInt(t *testing.T) {
	if ReturnInt() != 1 {
		t.Error("expected 1")
	}
}

func TestReturnFloat(t *testing.T) {
	if ReturnFloat() != float32(1.1) {
		t.Error("expected 1.1")
	}
}

func TestReturnIntArray(t *testing.T) {
	if ReturnIntArray() != [3]int{1, 3, 4} {
		t.Error("expected '[3]int{1, 3, 4}'")
	}
}

func TestReturnIntSlice(t *testing.T) {
	expected := []int{1, 2, 3}
	result := ReturnIntSlice()
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
}

func TestIntSliceToString(t *testing.T) {
	expected := "1723100500"
	result := IntSliceToString([]int{17, 23, 100500})
	if expected != result {
		t.Error("expected", expected, "have", result)
	}
}

func TestMergeSlices(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	result := MergeSlices([]float32{1.1, 2.1, 3.1}, []int32{4, 5})
	if !reflect.DeepEqual(result, expected) {
		t.Error("expected", expected, "have", result)
	}
}

func TestGetMapValuesSortedByKey(t *testing.T) {

	var cases = []struct {
		expected []string
		input    map[int]string
	}{
		{
			expected: []string{
				"Январь",
				"Февраль",
				"Март",
				"Апрель",
				"Май",
				"Июнь",
				"Июль",
				"Август",
				"Сентябрь",
				"Октябрь",
				"Ноябрь",
				"Декарь",
			},
			input: map[int]string{
				9:  "Сентябрь",
				1:  "Январь",
				2:  "Февраль",
				10: "Октябрь",
				5:  "Май",
				7:  "Июль",
				8:  "Август",
				12: "Декарь",
				3:  "Март",
				6:  "Июнь",
				4:  "Апрель",
				11: "Ноябрь",
			},
		},

		{
			expected: []string{
				"Зима",
				"Весна",
				"Лето",
				"Осень",
			},
			input: map[int]string{
				10: "Зима",
				30: "Лето",
				20: "Весна",
				40: "Осень",
			},
		},
	}

	for _, item := range cases {
		result := GetMapValuesSortedByKey(item.input)
		if !reflect.DeepEqual(result, item.expected) {
			t.Error("expected", item.expected, "have", result)
		}
	}
}
