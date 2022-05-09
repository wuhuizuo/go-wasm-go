package wazero

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// NewGoWASMStoreWithWazero prepare for wazero wasm runtime.
func NewGoWASMStoreWithWazero(b testing.TB, wasmFile string) (api.Module, func() error) {
	ctx := context.Background()

	source, err := os.ReadFile(wasmFile)
	if err != nil {
		b.Fatal(err)
	}

	runtime := wazero.NewRuntime()

	host, err := instantiateHostModuleForGo(ctx, runtime)
	if err != nil {
		b.Fatal(err)
	}

	compiled, err := runtime.CompileModule(ctx, source, wazero.NewCompileConfig())
	if err != nil {
		_ = host.Close(ctx)
		b.Fatal(err)
	}

	config := wazero.NewModuleConfig().WithName(wazeroModName)
	module, err := runtime.InstantiateModule(ctx, compiled, config)
	if err != nil {
		_ = host.Close(ctx)
		b.Fatal(err)
	}

	return module, func() (err error) {
		if e := module.Close(ctx); e != nil {
			err = e
		}
		if e := host.Close(ctx); e != nil && err != nil {
			err = e
		}
		return
	}
}

// CallGoWASMFuncWithWazero call test func with wazero loader.
func CallGoWASMFuncWithWazero(t testing.TB, module api.Module, funcName string, args ...uint64) []uint64 {
	f := module.ExportedFunction(funcName)
	if f == nil {
		t.Fatalf("not found func %s", funcName)
	}
	ret, err := module.ExportedFunction(funcName).Call(context.Background(), args...)
	if err != nil {
		t.Fatal(err)
	}
	if len(ret) != 1 {
		t.Fatalf("got values length is %d, but want %d ", len(ret), 1)
	}
	return ret
}

func instantiateHostModuleForGo(ctx context.Context, runtime wazero.Runtime) (api.Module, error) {
	return runtime.NewModuleBuilder("go").
		ExportFunctions(map[string]interface{}{
			"debug":                       func(sp int32) { fmt.Println(sp) },
			"runtime.resetMemoryDataView": func(int32) {},
			"runtime.wasmExit": func(ctx context.Context, m api.Module, code uint32) {
				_ = m.CloseWithExitCode(ctx, code)
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
		}).Instantiate(ctx)
}
