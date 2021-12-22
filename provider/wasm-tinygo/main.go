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
func RequestHTTP(in int32) int32 {
	return native.RequestHTTP()
}

//export FileIO
func FileIO(in int32) int32 {
	return native.FileIO()
}

//export MultiThreads
func MultiThreads(num int32) int32 {
	return native.MultiThreads(num)
}
