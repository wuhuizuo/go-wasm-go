package main

import (
	"os"
	"testing"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func getWasmedgeInstance(t testing.TB, wasmFile string) (*wasmedge.VM, *wasmedge.Configure) {
	conf := wasmedge.NewConfigure(wasmedge.WASI)
	vm := wasmedge.NewVMWithConfig(conf)
	wasi := vm.GetImportObject(wasmedge.HostRegistration(wasmedge.WASI))
	wasi.InitWasi(
		[]string{wasmFile}, /// The args
		os.Environ(),       /// The envs
		[]string{},         /// The preopens will be empty
	)

	/// Instantiate wasm
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

// callWASMFuncWithWasmedge call test func with wasmedge loaded func.
func callWASMFuncWithWasmedge(t testing.TB, vm *wasmedge.VM, in uint32) uint32 {
	ret, err := vm.ExecuteBindgen(wasmFuncName, wasmedge.Bindgen_return_i32, in)
	if err != nil {
		t.Fatal(err)
	}

	v, ok := ret.(int32)
	if !ok {
		t.Fatalf("return type is %T", ret)
	}

	return uint32(v)
}
