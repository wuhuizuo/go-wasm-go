package wasmtime

import (
	"fmt"
	"testing"

	"github.com/bytecodealliance/wasmtime-go"
)

// GetWasmFuncWithWasmtime get wasm func with wasmtime.
func GetGoWasmFuncWithWasmtime(t testing.TB, wasmFile, funcName string) (*wasmtime.Store, *wasmtime.Func) {
	store := newWasmStore()
	instance := newWasmInstance(store, wasmFile)

	// After we've instantiated we can lookup our `run` function and call
	// it.
	fn := instance.GetExport(store, funcName).Func()
	if fn == nil {
		panic("no exported func: " + funcName)
	}

	return store, fn
}

func newWasmInstance(store *wasmtime.Store, file string) *wasmtime.Instance {
	module, err := wasmtime.NewModuleFromFile(store.Engine, file)
	check(err)

	linker := newGoLinker(store)
	instance, err := linker.Instantiate(store, module)
	check(err)

	return instance
}

func newWasmStore() *wasmtime.Store {
	config := wasmtime.NewConfig()
	config.SetWasmModuleLinking(true)
	// config.SetInterruptable(true)
	// config.SetConsumeFuel(true)
	// config.SetDebugInfo(true)

	engine := wasmtime.NewEngineWithConfig(config)
	store := wasmtime.NewStore(engine)

	// wasi setting.
	// wasiConfig := wasmtime.NewWasiConfig()
	// wasiConfig.InheritEnv()
	// wasiConfig.SetEnv([]string{"WASMTIME"}, []string{"GO"})
	// store.SetWasi(wasiConfig)

	return store
}

func newGoLinker(store *wasmtime.Store) *wasmtime.Linker {
	linker := wasmtime.NewLinker(store.Engine)

	linker.DefineFunc(store, "go", "debug", func(ptr int32) { fmt.Sprintln(ptr) })
	linker.DefineFunc(store, "go", "runtime.resetMemoryDataView", func(int32) {})
	linker.DefineFunc(store, "go", "runtime.wasmExit", func(code int32) { panic(code) })
	linker.DefineFunc(store, "go", "runtime.wasmWrite", func(int32) {})
	linker.DefineFunc(store, "go", "runtime.nanotime1", func(int32) {})
	linker.DefineFunc(store, "go", "runtime.walltime", func(int32) {})
	linker.DefineFunc(store, "go", "runtime.scheduleTimeoutEvent", func(int32) {})
	linker.DefineFunc(store, "go", "runtime.clearTimeoutEvent", func(int32) {})
	linker.DefineFunc(store, "go", "runtime.getRandomData", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.finalizeRef", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.stringVal", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueGet", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueSet", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueIndex", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueSetIndex", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueInvoke", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueCall", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueNew", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueLength", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valuePrepareString", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueLoadString", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.valueInstanceOf", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.copyBytesToGo", func(int32) {})
	linker.DefineFunc(store, "go", "syscall/js.copyBytesToJS", func(int32) {})

	return linker
}
