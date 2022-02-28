package wasmtime

import (
	"testing"

	"github.com/bytecodealliance/wasmtime-go"
)

// GetWasmFuncWithWasmtime get wasm func with wasmtime.
func GetWasmFuncWithWasmtime(t testing.TB, wasmFile, funcName string) (*wasmtime.Store, *wasmtime.Func) {
	store := newWasiStore()
	instance := newWasiInstance(store, wasmFile)

	// After we've instantiated we can lookup our `run` function and call
	// it.
	fn := instance.GetExport(store, funcName).Func()
	if fn == nil {
		panic("no exported func: " + funcName)
	}

	return store, fn
}

func newWasiInstance(store *wasmtime.Store, file string) *wasmtime.Instance {
	module, err := wasmtime.NewModuleFromFile(store.Engine, file)
	check(err)

	linker := wasmtime.NewLinker(store.Engine)
	check(linker.DefineWasi())
	instance, err := linker.Instantiate(store, module)
	check(err)

	return instance
}

func newWasiStore() *wasmtime.Store {
	config := wasmtime.NewConfig()
	config.SetInterruptable(true)

	engine := wasmtime.NewEngineWithConfig(config)
	store := wasmtime.NewStore(engine)

	// wasi setting.
	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.InheritEnv()
	wasiConfig.SetEnv([]string{"WASMTIME"}, []string{"GO"})
	store.SetWasi(wasiConfig)

	return store
}
