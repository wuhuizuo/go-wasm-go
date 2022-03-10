package wasmtime

import (
	"testing"

	"github.com/bytecodealliance/wasmtime-go"
)

// GetRuntimes get wasm runtime with wasmtime.
func GetRuntimes(t testing.TB, wasmFile string, deps map[string]string) (*wasmtime.Store, *wasmtime.Instance) {
	store := newWasiStore()
	instance := newWasiInstance(store, wasmFile, deps)
	return store, instance
}

func CallFunc(t testing.TB, store *wasmtime.Store, instance *wasmtime.Instance, funcName string, args ...interface{}) (interface{}, error) {
	fn := instance.GetFunc(store, funcName)
	if fn == nil {
		panic("no exported func: " + funcName)
	}

	return fn.Call(store, args...)
}

func newWasiInstance(store *wasmtime.Store, file string, deps map[string]string) *wasmtime.Instance {
	depModules := make(map[string]*wasmtime.Module)
	for n, f := range deps {
		m, err := wasmtime.NewModuleFromFile(store.Engine, f)
		if err != nil {
			return nil
		}

		depModules[n] = m
	}

	module, err := wasmtime.NewModuleFromFile(store.Engine, file)
	check(err)

	linker := wasmtime.NewLinker(store.Engine)
	linker.AllowShadowing(true)
	check(linker.DefineWasi())

	// import dependent modules
	for n, m := range depModules {
		check(linker.DefineModule(store, n, m))
	}

	instance, err := linker.Instantiate(store, module)
	check(err)

	return instance
}

func newWasiStore() *wasmtime.Store {
	config := wasmtime.NewConfig()
	config.SetInterruptable(true)
	config.SetWasmReferenceTypes(true)
	config.SetWasmThreads(true)
	config.SetWasmBulkMemory(true)
	config.SetWasmModuleLinking(true)
	config.SetWasmMultiMemory(true)

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
