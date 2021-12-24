package main

import (
	"syscall/js"

	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

func main() {
	js.Global().Set("Fibonacci", js.FuncOf(Wrap(native.Fibonacci)))
	js.Global().Set("RequestHTTP", js.FuncOf(Wrap(native.RequestHTTP)))
	js.Global().Set("FileIO", js.FuncOf(Wrap(native.FileIO)))
	js.Global().Set("MultiThreads", js.FuncOf(Wrap(native.MultiThreads)))

	select {}
}
