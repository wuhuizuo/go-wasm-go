package main

import "github.com/wuhuizuo/go-wasm-go/provider/native"

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
func BytesTest(in []byte) []byte {
	return native.BytesTest(in)
}

//export InterfaceTest
func InterfaceTest(in interface{}) interface{} {
	return native.InterfaceTest(in)
}

//export ErrTest
func ErrTest(in error) error {
	return native.ErrTest(in)
}
