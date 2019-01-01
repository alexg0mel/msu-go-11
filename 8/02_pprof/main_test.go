package main

import (
	"regexp"
	"strings"
	"testing"
)

var haystack = `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras accumsan nisl et iaculis fringilla. Integer sapien orci, facilisis ut venenatis nec, suscipit at massa. Cras suscipit lectus non neque molestie, et imperdiet sem ultricies. Donec sit amet mattis nisi, efficitur posuere enim. Aliquam erat volutpat. Curabitur mattis nunc nisi, eu maximus dui facilisis in. Quisque vel tortor mauris. Praesent tellus sapien, vestibulum nec purus ut, luctus egestas odio. Ut ac ipsum non ipsum elementum pretium in id enim. Aenean eu augue fringilla, molestie orci et, tincidunt ipsum.
Nullam maximus odio vitae augue fermentum laoreet eget scelerisque ligula. Praesent pretium eu lacus in ornare. Maecenas fermentum id sapien non faucibus. Donec est tellus, auctor eu iaculis quis, accumsan vitae ligula. Fusce dolor nisl, pharetra eu facilisis non, hendrerit ac turpis. Pellentesque imperdiet aliquam quam in luctus. Curabitur ut orci sodales, faucibus nunc ac, maximus odio. Vivamus vitae nulla posuere, pellentesque quam posuere`

// go test -bench=. -cpuprofile cpu.out
// go tool pprof <бинарник(02_pprof.test)> cpu.out

func BenchmarkSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Contains(haystack, "auctor")
	}
}

func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexp.MatchString("auctor", haystack)
	}
}

//// прекомпиляция
//var reg = regexp.MustCompile("auctor")
//
//func BenchmarkRegex2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		reg.MatchString(haystack)
//	}
//}
//
//// тестим сгенерированную регулярку github.com/opennota/re2dfa
//// re2dfa auctor main.matchAuctor string
//func BenchmarkGeneratedRegex(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		matchAuctor(haystack)
//	}
//}
