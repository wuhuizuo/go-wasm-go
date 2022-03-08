package main

//go:generate tinygo build -target wasi -wasm-abi=generic -o wasi.wasm

func main() {
	// nothing.
}

//go:module standalone
//export Fibonacci
func Fibonacci(in int32) int32

//go:module standalone
//export BytesTest
func BytesTest(in []byte) int64

//export RunInt32
func RunInt32(in int32) int32 {
	return Fibonacci(in)
}

//export Run
func RunBytes(in []byte) int64 {
	return BytesTest(in)
}
