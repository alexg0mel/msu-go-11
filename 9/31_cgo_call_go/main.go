package main

/*
https://golang.org/cmd/cgo/#hdr-C_references_to_Go
см. послений абзац

main.cgo2.o: In function `Multiply':
9/31_cgo_call_go/main.go:10: multiple definition of `Multiply'
9/31_cgo_call_go
/main.go:10: first defined here
collect2.exe: error: ld returned 1 exit status
*/

/*
void Multiply(int a, int b);
*/
import "C" //это псевдо-пакет, он реализуется компилятором
import "fmt"

//export printResult
func printResult(result C.int) {
	fmt.Printf("result-var internals %T = %+v\n", result, result)
}

func main() {
	a := 2
	b := 3
	C.Multiply(C.int(a), C.int(b))
}
