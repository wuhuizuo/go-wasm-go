package main

import (
	"unsafe"

	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

var buffer []byte

func main() {
	// nothing.
}

//go:wasmexport
func Fibonacci(in int32) int32 {
	return native.Fibonacci(in)
}

//go:wasmexport
func RequestHTTP() int32 {
	return native.RequestHTTP()
}

//go:wasmexport
func FileIO() int32 {
	return native.FileIO()
}

//go:wasmexport
func MultiThreads(num int32) int32 {
	return native.MultiThreads(num)
}

//go:wasmexport
func BytesTest(in []byte) int32 {
	buffer = native.BytesTest(in)
	return *(*int32)(unsafe.Pointer(&buffer))
}

//go:wasmexport
func BytesTestLen() int32 {
	return int32(len(buffer))
}

//go:wasmexport
func InterfaceTest(in interface{}) interface{} {
	return native.InterfaceTest(in)
}

//go:wasmexport
func ErrTest(in error) error {
	return native.ErrTest(in)
}
