package main

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	epsilon := 0.01
	for i, tt := range []float64{-1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		res := Sqrt(tt)
		expected := math.Sqrt(tt)
		if math.Abs(expected-res) > epsilon {
			t.Errorf("Case[%d]: bad result, expected %v, got %v", i, expected, res)
		}
	}
}
