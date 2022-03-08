package wazero

import (
	"os"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasm"
)

const wazeroModName = "wasmtest"

// NewWASMStoreWithWazero prepare for wazero wasm store.
func NewWASMStoreWithWazero(b testing.TB, wasmFile string) wasm.ModuleExports {
	binary, err := os.ReadFile(wasmFile)
	if err != nil {
		b.Fatal(err)
	}

	store := wazero.NewStoreWithConfig(&wazero.StoreConfig{Engine: wazero.NewEngineJIT()})

	_, err = wazero.InstantiateHostModule(store, wazero.WASISnapshotPreview1())
	if err != nil {
		b.Fatal(err)
	}

	exports, err := wazero.InstantiateModule(store, &wazero.ModuleConfig{
		Name:   wazeroModName,
		Source: binary,
	})
	if err != nil {
		b.Fatal(err)
	}

	return exports
}

// CallWASMFuncWithWazero call test func with wazero loader.
func CallWASMFuncWithWazero(t testing.TB, exports wasm.ModuleExports, funcName string, args ...uint64) []uint64 {
	ret, err := exports.Function(funcName).Call(nil, args...)
	if err != nil {
		t.Fatal(err)
	}
	if len(ret) != 1 {
		t.Fatalf("got values length is %d, but want %d ", len(ret), 1)
	}
	return ret
}
