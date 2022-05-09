package wasmtime

import (
	"testing"

	"github.com/bytecodealliance/wasmtime-go"
)

// GetWasmFuncWithWasmtime get wasm func with wasmtime.
func GetWasmFuncWithWasmtime(t testing.TB, wasmFile string) (*wasmtime.Store, *wasmtime.Instance) {
	store := newWasiStore()
	instance := newWasiInstance(store, wasmFile)
	return store, instance
}

func CallWasmFunc(t testing.TB, store *wasmtime.Store, instance *wasmtime.Instance, funcName string, args ...interface{}) (interface{}, error) {
	fn := instance.GetFunc(store, funcName)
	if fn == nil {
		panic("no exported func: " + funcName)
	}

	return fn.Call(store, args...)
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
	config.SetWasmReferenceTypes(true)
	config.SetWasmThreads(true)
	config.SetWasmBulkMemory(true)

	engine := wasmtime.NewEngineWithConfig(config)
	store := wasmtime.NewStore(engine)

	// wasi setting.
	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.InheritStdout()
	wasiConfig.InheritStderr()
	wasiConfig.InheritStdin()
	wasiConfig.InheritEnv()
	wasiConfig.SetEnv([]string{"WASMTIME"}, []string{"GO"})
	store.SetWasi(wasiConfig)

	return store
}
