package main

import (
	"bytes"
	"fmt"
	"testing"
)

func BenchmarkPrint(b *testing.B) {
	var buf bytes.Buffer
	var s string = "test string"
	for i := 0; i < b.N; i++ {
		buf.Reset()
		fmt.Fprintf(&buf, "string is: %s", s)
		//buf.Write([]byte("string is: " + s))
	}
}
