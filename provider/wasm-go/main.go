package main

import (
	"syscall/js"

	"github.com/mattn/gowasmer/wasmutil"
	"github.com/wuhuizuo/go-wasm-go/provider"
)

func main() {
	js.Global().Set("Fibonacci", js.FuncOf(wasmutil.Wrap(provider.Fibonacci)))
	js.Global().Set("HTTPBasicAuth", js.FuncOf(wasmutil.Wrap(provider.HTTPBasicAuth)))

	select {}
}
