package main

import (
	"io/ioutil"
	"testing"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

// getWasmFuncWithWasmer parse wasm function with wasmer.
func getWasmFuncWithWasmer(t testing.TB, wasmFile string) func(...interface{}) (interface{}, error) {
	binary, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		t.Fatal(err)
	}

	store := wasmer.NewStore(wasmer.NewEngine())
	// defer store.Close()

	// Compiles the mod
	mod, err := wasmer.NewModule(store, binary)
	if err != nil {
		t.Fatal(err)
	}
	// defer mod.Close()

	// wasi dealing.
	wasiEnv, _ := wasmer.NewWasiStateBuilder("wasi-program").
		// Choose according to your actual situation
		// Argument("--foo").
		// Environment("ABC", "DEF").
		// MapDirectory("./", ".").
		Finalize()

	// Instantiates the module
	importObject, err := wasiEnv.GenerateImportObject(store, mod)
	if err != nil {
		t.Fatal(err)
	}

	instance, err := wasmer.NewInstance(mod, importObject)
	if err != nil {
		t.Fatal(err)
	}
	// defer instance.Close()

	// Gets the `fn` exported function from the WebAssembly instance.
	fn, err := instance.Exports.GetFunction(wasmFuncName)
	if err != nil {
		t.Fatal(err)
	}

	return fn
}

// callWASMFuncWithWasmer call test func with wasmer loaded func.
func callWASMFuncWithWasmer(t testing.TB, fn func(...interface{}) (interface{}, error), in uint32) uint32 {
	// 这里有点特殊, uint系列会被转换成 int系列
	ret, err := fn(int32(in))
	if err != nil {
		t.Fatal(err)
	}

	v, ok := ret.(int32)
	if !ok {
		t.Fatalf("return type is %T", ret)
	}

	return uint32(v)
}
