package main

//go:generate tinygo build -target wasi -wasm-abi=generic -o wasi.wasm

func main() {
	// nothing.
}

//go:wasm-module standalone
//export Fibonacci
func _Fibonacci(in int32) int32

//go:wasm-module standalone
//export BytesTest
func _BytesTest(in []byte) int64

//go:wasm-module standalone
//export StringTest
func _StringTest(in string) int64

//export Fibonacci
func Fibonacci(in int32) int32 {
	return _Fibonacci(in)
}

//export BytesTest
func BytesTest(in []byte) int64 {
	return _BytesTest(in)
}

//export StringTest
func StringTest(in string) int64 {
	println("-p--|", in, "|---")
	return _StringTest(in)
}
