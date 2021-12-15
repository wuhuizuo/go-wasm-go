package main

import "github.com/wuhuizuo/go-wasm-go/provider"

func main() {
	// nothing.
}

//export Fibonacci
func Fibonacci(in int32) int32 {
	return provider.Fibonacci(in)
}

//export HTTPBasicAuth
func HTTPBasicAuth(username, password string) {
	provider.HTTPBasicAuth(username, password)
}
