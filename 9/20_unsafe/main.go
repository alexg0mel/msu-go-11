package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"unsafe"
)

type Message struct {
	flag1 bool
	name  string
	flag2 bool
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func main() {
	a := int(1)
	fmt.Println("memory pointer for a", unsafe.Pointer(&a))
	fmt.Println("memory size for a", unsafe.Sizeof(a))

	println("-------")

	msg := Message{flag1: false, flag2: false, name: "new message"}
	fmt.Println("memory size for Message struct", unsafe.Sizeof(msg))
	fmt.Println("flag1 Sizeof", unsafe.Sizeof(msg.flag1), "Alignof", unsafe.Alignof(msg.flag1), "Offsetof", unsafe.Offsetof(msg.flag1))
	fmt.Println("name Sizeof", unsafe.Sizeof(msg.name), "Alignof", unsafe.Alignof(msg.name), "Offsetof", unsafe.Offsetof(msg.name))
	fmt.Println("flag2 Sizeof", unsafe.Sizeof(msg.flag2), "Alignof", unsafe.Alignof(msg.flag2), "Offsetof", unsafe.Offsetof(msg.flag2))

	println("-------")

	fmt.Printf("%#016x\n", Float64bits(10.0))
	fmt.Printf("%b\n", Float64bits(1.0))
}
