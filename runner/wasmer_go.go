package runner

import (
	"io/ioutil"
	"testing"

	"github.com/mattn/gowasmer"
)

// getWasmFuncWithWasmer parse wasm function with wasmer.
func getGoWasmFuncWithWasmer(t testing.TB, wasmFile, funcName string) interface{} {
	binary, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		t.Fatal(err)
	}

	instance, err := gowasmer.NewInstance(binary)
	if err != nil {
		t.Fatal(err)
	}

	// Gets the `fn` exported function from the WebAssembly instance.
	fn := instance.Get(funcName)
	if fn == nil {
		t.Fatal("not found exported function in wasm exec binnay")
	}

	return fn
}

// callGoWASMFuncWithWasmer call test func with wasmer loaded func.
func callGoWASMFuncWithWasmer(t testing.TB, fn interface{}, args []interface{}) interface{} {
	switch v := fn.(type) {
	case func([]interface{}) interface{}:
		return v(args)
	default:
		t.Fatalf("unknown type: %T", fn)
		return nil
	}
}
