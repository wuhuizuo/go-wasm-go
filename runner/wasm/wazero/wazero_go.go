package wazero

import (
	"fmt"
	"os"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/wasm"
)

// NewGoWASMStoreWithWazero prepare for wazero wasm runtime.
func NewGoWASMStoreWithWazero(b testing.TB, wasmFile string) (wasm.Module, func() error) {
	binary, err := os.ReadFile(wasmFile)
	if err != nil {
		b.Fatal(err)
	}

	runtime := wazero.NewRuntime()

	host, err := instantiateHostModuleForGo(runtime)
	if err != nil {
		b.Fatal(err)
	}

	decoded, err := runtime.CompileModule(binary)
	if err != nil {
		host.Close()
		b.Fatal(err)
	}

	module, err := runtime.InstantiateModule(decoded.WithName(wazeroModName))
	if err != nil {
		host.Close()
		b.Fatal(err)
	}

	return module, func() (err error) {
		module.Close()
		host.Close()
		return
	}
}

// CallGoWASMFuncWithWazero call test func with wazero loader.
func CallGoWASMFuncWithWazero(t testing.TB, module wasm.Module, funcName string, args ...uint64) []uint64 {
	f := module.ExportedFunction(funcName)
	if f == nil {
		t.Fatalf("not found func %s", funcName)
	}
	ret, err := module.ExportedFunction(funcName).Call(nil, args...)
	if err != nil {
		t.Fatal(err)
	}
	if len(ret) != 1 {
		t.Fatalf("got values length is %d, but want %d ", len(ret), 1)
	}
	return ret
}

func instantiateHostModuleForGo(runtime wazero.Runtime) (wasm.Module, error) {
	return runtime.NewModuleBuilder("go").
		ExportFunctions(map[string]interface{}{
			"debug":                         func(sp int32) { fmt.Println(sp) },
			"runtime.resetMemoryDataView":   func(int32) {},
			"runtime.wasmExit":              func(m wasm.Module, code uint32) {
				_ = m.CloseWithExitCode(code)
			},
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
		}).Instantiate()
}
