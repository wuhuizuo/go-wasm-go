package wasmer

import (
	"fmt"
	"io/ioutil"
	"runtime/debug"
	"testing"
)

// GetGoWasmFuncWithWasmer parse wasm function with wasmer.
func GetGoWasmFuncWithWasmer(t testing.TB, wasmFile, funcName string) interface{} {
	debug.SetGCPercent(-1)
	binary, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("---------------")

	instance, err := NewInstance(binary)
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

// CallGoWASMFuncWithWasmer call test func with wasmer loaded func.
func CallGoWASMFuncWithWasmer(t testing.TB, fn interface{}, args []interface{}) interface{} {
	switch v := fn.(type) {
	case func([]interface{}) interface{}:
		return v(args)
	default:
		t.Fatalf("unknown type: %T", fn)
		return nil
	}
}
