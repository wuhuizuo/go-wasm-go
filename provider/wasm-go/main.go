package main

import (
	"syscall/js"

	"github.com/mattn/gowasmer/wasmutil"
	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

func main() {
	js.Global().Set("Fibonacci", js.FuncOf(wasmutil.Wrap(native.Fibonacci)))
	js.Global().Set("RequestHTTP", js.FuncOf(wasmutil.Wrap(native.RequestHTTP)))

	select {}
}
