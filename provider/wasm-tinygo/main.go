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
func RequestHTTP() {
	return
	// native.RequestHTTP()
}
