package main

import (
	"syscall/js"

	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

//go:generate env GOOS=js GOARCH=wasm go build -o wasm.wasm

func main() {
	js.Global().Set("Fibonacci", js.FuncOf(Wrap(native.Fibonacci)))
	js.Global().Set("RequestHTTP", js.FuncOf(Wrap(native.RequestHTTP)))
	js.Global().Set("FileIO", js.FuncOf(Wrap(native.FileIO)))
	js.Global().Set("MultiThreads", js.FuncOf(Wrap(native.MultiThreads)))
	js.Global().Set("BytesTest", js.FuncOf(Wrap(native.BytesTest)))
	js.Global().Set("StringTest", js.FuncOf(Wrap(native.StringTest)))
	js.Global().Set("InterfaceTest", js.FuncOf(Wrap(native.InterfaceTest)))
	js.Global().Set("ErrTest", js.FuncOf(Wrap(native.ErrTest)))

	select {}
}
