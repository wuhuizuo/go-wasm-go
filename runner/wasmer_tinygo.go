package runner

import (
	"io/ioutil"
	"testing"

	"github.com/wasmerio/wasmer-go/wasmer"
)

// getWasmFuncWithWasmer parse wasm function with wasmer.
func getWasmFuncWithWasmer(t testing.TB, wasmFile, funcName string) func(...interface{}) (interface{}, error) {
	binary, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		t.Fatal(err)
	}

	store := wasmer.NewStore(wasmer.NewEngine())

	// Compiles the mod
	mod, err := wasmer.NewModule(store, binary)
	if err != nil {
		t.Fatal(err)
	}

	// wasi dealing.
	wasiEnv, _ := wasmer.NewWasiStateBuilder("wasi-program").Finalize()

	// Instantiates the module
	importObject, err := wasiEnv.GenerateImportObject(store, mod)
	if err != nil {
		t.Fatal(err)
	}

	instance, err := wasmer.NewInstance(mod, importObject)
	if err != nil {
		t.Fatal(err)
	}

	// Gets the `fn` exported function from the WebAssembly instance.
	fn, err := instance.Exports.GetFunction(funcName)
	if err != nil {
		t.Fatal(err)
	}

	return fn
}

// callWASMFuncWithWasmer call test func with wasmer loaded func.
func callWASMFuncWithWasmer(t testing.TB, fn func(...interface{}) (interface{}, error), args ...interface{}) interface{} {
	// 这里有点特殊, uint系列会被转换成 int系列
	ret, err := fn(args...)
	if err != nil {
		t.Fatal(err)
	}

	return ret
}
