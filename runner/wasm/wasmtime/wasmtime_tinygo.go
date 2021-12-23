package wasmtime

import (
	"io/ioutil"
	"testing"

	"github.com/bytecodealliance/wasmtime-go"
)

// GetWasmFuncWithWasmtime get wasm func with wasmtime.
func GetWasmFuncWithWasmtime(t testing.TB, wasmFile, funcName string) (*wasmtime.Store, *wasmtime.Func) {
	binary, err := ioutil.ReadFile(wasmFile)
	check(err)

	cc := wasmtime.NewConfig()
	cc.SetInterruptable(true)
	// Almost all operations in wasmtime require a contextual `store`
	// argument to share, so create that first
	store := wasmtime.NewStore(wasmtime.NewEngineWithConfig(cc))
	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.SetEnv([]string{"WASMTIME"}, []string{"GO"})
	store.SetWasi(wasiConfig)

	// Once we have our binary `wasm` we can compile that into a `*Module`
	// which represents compiled JIT code.
	module, err := wasmtime.NewModule(store.Engine, binary)
	check(err)
	module.AsExtern()

	// Next up we instantiate a module which is where we link in all our
	// imports. We've got one import so we pass that in here.
	instance, err := wasmtime.NewInstance(store, module, nil)
	check(err)

	// After we've instantiated we can lookup our `run` function and call
	// it.
	fn := instance.GetExport(store, funcName).Func()
	if fn == nil {
		panic("no exported func: " + funcName)
	}

	return store, fn
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
