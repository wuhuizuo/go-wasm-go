package runner

import (
	"io/ioutil"
	"testing"

	"github.com/mattn/gowasmer"
)

// getWasmFuncWithWasmer parse wasm function with wasmer.
func getGoWasmFuncWithWasmer(t testing.TB, wasmFile string) interface{} {
	binary, err := ioutil.ReadFile(wasmFile)
	if err != nil {
		t.Fatal(err)
	}

	instance, err := gowasmer.NewInstance(binary)
	if err != nil {
		t.Fatal(err)
	}

	// Gets the `fn` exported function from the WebAssembly instance.
	fn := instance.Get(fibFuncName)
	if fn == nil {
		t.Fatal("not found exported function in wasm exec binnay")
	}

	return fn
}

// callWASMFuncWithWasmer call test func with wasmer loaded func.
func callGoWASMFuncWithWasmer(t testing.TB, fn interface{}, in int32) int32 {
	switch v := fn.(type) {
	case func([]interface{}) interface{}:
		ret := v([]interface{}{int32(in)})

		reti, ok := ret.(float64)
		if !ok {
			t.Fatalf("return type is %T", ret)
		}

		return int32(reti)
	default:
		t.Fatalf("unknown type: %T", fn)
		return 0
	}
}
