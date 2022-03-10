package main

import (
	"unsafe"

	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

//go:generate tinygo build -target wasi -wasm-abi=generic -o wasi.wasm

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

// BytesTest return int64, first 4bytes present pointer, last 4bytes present data length.
//export BytesTest
func BytesTest(in []byte) int64 {
	buffer := native.BytesTest(in)
	ptr := *(*int32)(unsafe.Pointer(&buffer))

	return int64(ptr)<<32 | int64(len(buffer))
}

// StringTest return int64, first 4bytes present pointer, last 4bytes present data length.
//export StringTest
func StringTest(in string) int64 {
	buffer := native.StringTest(in)
	ptr := *(*int32)(unsafe.Pointer(&buffer))

	return int64(ptr)<<32 | int64(len(buffer))
}

//export InterfaceTest
func InterfaceTest(in interface{}) interface{} {
	return native.InterfaceTest(in)
}

//export ErrTest
func ErrTest(in error) error {
	return native.ErrTest(in)
}
