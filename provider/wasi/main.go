package main

import (
	"unsafe"

	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

//go:generate tinygo build -o wasi.wasm -target wasi

var buffer []byte

func main() {
	// nothing.
}

//export Fibonacci
func Fibonacci(in int32) int32 {
	return native.Fibonacci(in)
}

//export RequestHTTP
func RequestHTTP() int32 {
	return native.RequestHTTP()
}

//export FileIO
func FileIO() int32 {
	return native.FileIO()
}

//export MultiThreads
func MultiThreads(num int32) int32 {
	return native.MultiThreads(num)
}

//export BytesTest
func BytesTest(in []byte) int32 {
	buffer = native.BytesTest(in)
	return *(*int32)(unsafe.Pointer(&buffer))
}

//export BytesTestLen
func BytesTestLen() int32 {
	return int32(len(buffer))
}

//export InterfaceTest
func InterfaceTest(in interface{}) interface{} {
	return native.InterfaceTest(in)
}

//export ErrTest
func ErrTest(in error) error {
	return native.ErrTest(in)
}
