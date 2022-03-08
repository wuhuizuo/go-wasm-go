package runner

import (
	"fmt"
	"path/filepath"
	"testing"

	wg "github.com/bytecodealliance/wasmtime-go"
	"github.com/stretchr/testify/assert"

	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmtime"
)

func Test_wasmtime_tinygo(t *testing.T) {
	store, instance := wasmtime.GetWasmFuncWithWasmtime(t, filepath.Join(selfDir(t), "..", wasi))

	t.Run("algorithm", func(t *testing.T) {
		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := wasmtime.CallWasmFunc(t, store, instance, fibFuncName, tt.in)

				assert.NoError(t, err)
				assert.IsType(t, tt.want, got)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("bytes test", func(t *testing.T) {
		inPtr, inSize, inCap := wasmtime.TransInBytesParam(store, instance, []byte("hello"))
		got, err := wasmtime.CallWasmFunc(t, store, instance, byteInOutFuncName, inPtr, inSize, inCap)
		t.Log(got, err)
		assert.NoError(t, err)

		outSize, err := wasmtime.CallWasmFunc(t, store, instance, byteInOutLenFuncName)
		t.Log(outSize, err)
		assert.NoError(t, err)

		out := wasmtime.ReadOutBytesReturn(store, instance, got.(int32), outSize.(int32))
		assert.Equal(t, []byte("hello---"), out)
	})
}

func Test_wasmtime_go(t *testing.T) {
	t.Skip()
	t.Run("algorithm", func(t *testing.T) {
		store, fn := wasmtime.GetGoWasmFuncWithWasmtime(t, filepath.Join(selfDir(t), "..", wasmGo), fibFuncName)

		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := fn.Call(store, tt.in)

				assert.NoError(t, err)
				assert.IsType(t, tt.want, got)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}

func ExampleLinkModule() {
	store := wg.NewStore(wg.NewEngine())

	// Compile two wasm modules where the first references the second
	wasm1, err := wg.Wat2Wasm(`
	    (module
		(import "wasm2" "double" (func $double (param i32) (result i32)))
		(func (export "double_and_add") (param i32 i32) (result i32)
		  local.get 0
		  call $double
		  local.get 1
		  i32.add)
	    )
	`)
	checkerr(err)

	wasm2, err := wg.Wat2Wasm(`
	    (module
		(func (export "double") (param i32) (result i32)
		  local.get 0
		  i32.const 2
		  i32.mul)
	    )
	`)
	checkerr(err)

	// Next compile both modules
	module1, err := wg.NewModule(store.Engine, wasm1)
	checkerr(err)
	module2, err := wg.NewModule(store.Engine, wasm2)
	checkerr(err)

	linker := wg.NewLinker(store.Engine)
	err = linker.DefineModule(store, "wasm2", module2)
	checkerr(err)

	instance1, err := linker.Instantiate(store, module1)
	checkerr(err)
	doubleAndAdd := instance1.GetFunc(store, "double_and_add")
	result, err := doubleAndAdd.Call(store, 2, 3)
	checkerr(err)
	fmt.Print(result.(int32))
	// Output: 7
}

func ExampleLinker() {
	store := wg.NewStore(wg.NewEngine())

	// Compile two wasm modules where the first references the second
	wasm1, err := wg.Wat2Wasm(`
	    (module
		(import "wasm2" "double" (func $double (param i32) (result i32)))
		(func (export "double_and_add") (param i32 i32) (result i32)
		  local.get 0
		  call $double
		  local.get 1
		  i32.add)
	    )
	`)
	checkerr(err)

	wasm2, err := wg.Wat2Wasm(`
	    (module
		(func (export "double") (param i32) (result i32)
		  local.get 0
		  i32.const 2
		  i32.mul)
	    )
	`)
	checkerr(err)

	// Next compile both modules
	module1, err := wg.NewModule(store.Engine, wasm1)
	checkerr(err)
	module2, err := wg.NewModule(store.Engine, wasm2)
	checkerr(err)

	linker := wg.NewLinker(store.Engine)

	// The second module is instantiated first since it has no imports, and
	// then we insert the instance back into the linker under the name
	// the first module expects.
	instance2, err := linker.Instantiate(store, module2)
	checkerr(err)
	err = linker.DefineInstance(store, "wasm2", instance2)
	checkerr(err)

	// And now we can instantiate our first module, executing the result
	// afterwards
	instance1, err := linker.Instantiate(store, module1)
	checkerr(err)
	doubleAndAdd := instance1.GetFunc(store, "double_and_add")
	result, err := doubleAndAdd.Call(store, 2, 3)
	checkerr(err)
	fmt.Print(result.(int32))
	// Output: 7
}

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
