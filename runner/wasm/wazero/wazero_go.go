package wazero

import (
	"fmt"
	"os"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasm"
)

// NewGoWASMStoreWithWazero prepare for wazero wasm store.
func NewGoWASMStoreWithWazero(b testing.TB, wasmFile string) wasm.ModuleExports {
	binary, err := os.ReadFile(wasmFile)
	if err != nil {
		b.Fatal(err)
	}

	store := wazero.NewStoreWithConfig(&wazero.StoreConfig{Engine: wazero.NewEngineJIT()})

	if err := instantiateHostModuleForGo(store); err != nil {
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

// CallGoWASMFuncWithWazero call test func with wazero loader.
func CallGoWASMFuncWithWazero(t testing.TB, exports wasm.ModuleExports, funcName string, args ...uint64) []uint64 {
	f := exports.Function(funcName)
	if f == nil {
		t.Fatalf("not found func %s", funcName)
	}
	ret, err := exports.Function(funcName).Call(nil, args...)
	if err != nil {
		t.Fatal(err)
	}
	if len(ret) != 1 {
		t.Fatalf("got values length is %d, but want %d ", len(ret), 1)
	}
	return ret
}

func instantiateHostModuleForGo(store wasm.Store) error {
	wazero.WASISnapshotPreview1()
	_, err := wazero.InstantiateHostModule(store, &wazero.HostModuleConfig{
		Name: "go",
		Functions: map[string]interface{}{
			"debug":                         func(sp int32) { fmt.Println(sp) },
			"runtime.resetMemoryDataView":   func(int32) {},
			"runtime.wasmExit":              func(code int32) { os.Exit(int(code)) },
			"runtime.wasmWrite":             func(int32) {},
			"runtime.nanotime1":             func(int32) {},
			"runtime.walltime":              func(int32) {},
			"runtime.scheduleTimeoutEvent":  func(int32) {},
			"runtime.clearTimeoutEvent":     func(int32) {},
			"runtime.getRandomData":         func(int32) {},
			"syscall/js.finalizeRef":        func(int32) {},
			"syscall/js.stringVal":          func(int32) {},
			"syscall/js.valueGet":           func(int32) {},
			"syscall/js.valueSet":           func(int32) {},
			"syscall/js.valueIndex":         func(int32) {},
			"syscall/js.valueSetIndex":      func(int32) {},
			"syscall/js.valueInvoke":        func(int32) {},
			"syscall/js.valueCall":          func(int32) {},
			"syscall/js.valueNew":           func(int32) {},
			"syscall/js.valueLength":        func(int32) {},
			"syscall/js.valuePrepareString": func(int32) {},
			"syscall/js.valueLoadString":    func(int32) {},
			"syscall/js.valueInstanceOf":    func(int32) {},
			"syscall/js.copyBytesToGo":      func(int32) {},
			"syscall/js.copyBytesToJS":      func(int32) {},
		},
	})

	return err
}
