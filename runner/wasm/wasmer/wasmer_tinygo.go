package wasmer

import (
	"io/ioutil"
	"testing"

	"github.com/wasmerio/wasmer-go/wasmer"
)

// GetWasmFuncWithWasmer parse wasm function with wasmer.
func GetWasmFuncWithWasmer(t testing.TB, wasmFile, funcName string) wasmer.NativeFunction {
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
	builder := wasmer.NewWasiStateBuilder("wasm-tinygo").
		Environment("KEY", "VALUE").
		CaptureStdout()
	wasiEnv, err := builder.Finalize()
	if err != nil {
		t.Fatal(err)
	}

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

// CallWASMFuncWithWasmer call test func with wasmer loaded func.
func CallWASMFuncWithWasmer(t testing.TB, fn wasmer.NativeFunction, args ...interface{}) interface{} {
	ret, err := fn(args...)
	if err != nil {
		t.Fatal(err)
	}

	return ret
}
