package main

import (
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	for i, tt := range []interface{}{1, uint(2), int8(4), float64(7), "", 'w', []int{}, map[string]bool{}} {
		if showMeTheType(tt) != fmt.Sprintf("%T", tt) {
			t.Errorf("Case[%d]: failed res %s, expected %T", i, showMeTheType(tt), tt)
		}
	}
}
