package main

import (
	"unsafe"

	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

//go:generate tinygo build -target=wasm -wasm-abi=generic -o wasm.wasm

var buffer []byte

func main() {
	// nothing.
}

// OtherFunc import from other module.
//go:wasm-module other
//export other.OtherFunc
func OtherFunc(in int32) int32

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
	return OtherFunc(native.FileIO())
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
