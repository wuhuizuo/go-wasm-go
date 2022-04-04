package wazero

import (
	"os"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/wasi"
)

const wazeroModName = "wasmtest"

// NewWASMStoreWithWazero prepare for wazero wasm runtime.
func NewWASMStoreWithWazero(b testing.TB, wasmFile string) (api.Module, func() error) {
	binary, err := os.ReadFile(wasmFile)
	if err != nil {
		b.Fatal(err)
	}

	runtime := wazero.NewRuntime()

	wm, err := wasi.InstantiateSnapshotPreview1(runtime)
	if err != nil {
		wm.Close()
		b.Fatal(err)
	}

	code, err := runtime.CompileModule(binary)
	if err != nil {
		wm.Close()
		b.Fatal(err)
	}

	config := wazero.NewModuleConfig().WithName(wazeroModName)
	module, err := runtime.InstantiateModuleWithConfig(code, config)
	if err != nil {
		wm.Close()
		b.Fatal(err)
	}

	return module, func() (err error) {
		module.Close()
		wm.Close()
		return
	}
}

// CallWASMFuncWithWazero call test func with wazero loader.
func CallWASMFuncWithWazero(t testing.TB, module api.Module, funcName string, args ...uint64) []uint64 {
	ret, err := module.ExportedFunction(funcName).Call(nil, args...)
	if err != nil {
		t.Fatal(err)
	}
	if len(ret) != 1 {
		t.Fatalf("got values length is %d, but want %d ", len(ret), 1)
	}
	return ret
}
