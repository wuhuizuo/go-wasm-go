package wazero

import (
	"os"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasm"
)

const wazeroModName = "wasmtest"

// NewWASMStoreWithWazero prepare for wazero wasm runtime.
func NewWASMStoreWithWazero(b testing.TB, wasmFile string) (wasm.Module, func() error) {
	binary, err := os.ReadFile(wasmFile)
	if err != nil {
		b.Fatal(err)
	}

	runtime := wazero.NewRuntime()

	wasi, err := runtime.InstantiateModule(wazero.WASISnapshotPreview1())
	if err != nil {
		wasi.Close()
		b.Fatal(err)
	}

	decoded, err := runtime.CompileModule(binary)
	if err != nil {
		wasi.Close()
		b.Fatal(err)
	}

	module, err := runtime.InstantiateModule(decoded.WithName(wazeroModName))
	if err != nil {
		wasi.Close()
		b.Fatal(err)
	}

	return module, func() (err error) {
		module.Close()
		wasi.Close()
		return
	}
}

// CallWASMFuncWithWazero call test func with wazero loader.
func CallWASMFuncWithWazero(t testing.TB, module wasm.Module, funcName string, args ...uint64) []uint64 {
	ret, err := module.ExportedFunction(funcName).Call(nil, args...)
	if err != nil {
		t.Fatal(err)
	}
	if len(ret) != 1 {
		t.Fatalf("got values length is %d, but want %d ", len(ret), 1)
	}
	return ret
}
