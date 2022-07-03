package wasmedge

import (
	"os"
	"testing"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

// GetWasmedgeInstance get wasm instance with wasmedge.
func GetWasmedgeInstance(t testing.TB, wasmFile string) (*wasmedge.VM, *wasmedge.Configure) {
	conf := wasmedge.NewConfigure(wasmedge.WASI)
	vm := wasmedge.NewVMWithConfig(conf)
	wasi := vm.GetImportModule(wasmedge.HostRegistration(wasmedge.WASI))
	wasi.InitWasi(
		[]string{wasmFile}, /// The args
		os.Environ(),       /// The envs
		[]string{},         /// The preopens will be empty
	)

	// Instantiate wasm
	if err := vm.LoadWasmFile(wasmFile); err != nil {
		t.Fatal(err)
	}
	if err := vm.Validate(); err != nil {
		t.Fatal(err)
	}
	if err := vm.Instantiate(); err != nil {
		t.Fatal(err)
	}

	return vm, conf
}

// CallWASMFuncWithWasmedgeReturnInt32 call test func with wasmedge loaded func.
func CallWASMFuncWithWasmedgeReturnInt32(t testing.TB, vm *wasmedge.VM, funcName string, args ...interface{}) int32 {
	ret, err := vm.ExecuteBindgen(funcName, wasmedge.Bindgen_return_i32, args...)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := ret.(int32)
	if !ok {
		t.Fatalf("return type is %T", ret)
	}

	return v
}

func CallWASMFuncWithWasmedgeReturnVoid(t testing.TB, vm *wasmedge.VM, funcName string, args ...interface{}) {
	_, err := vm.ExecuteBindgen(funcName, wasmedge.Bindgen_return_void, args...)
	if err != nil {
		t.Fatal(err)
	}
}
