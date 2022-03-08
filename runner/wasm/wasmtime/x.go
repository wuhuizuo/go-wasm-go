package wasmtime

import (
	"github.com/bytecodealliance/wasmtime-go"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TransInBytesParam(store *wasmtime.Store, instance *wasmtime.Instance, in []byte) (_, _, _ interface{}) {
	// malloc memory space.
	calloc := instance.GetFunc(store, "malloc")
	aret, _ := calloc.Call(store, len(in))
	start := int(aret.(int32))

	// write to wasm vm memory.
	bs := instance.GetExport(store, "memory").Memory().UnsafeData(store)
	for i, c := range in {
		bs[start+i] = c
	}

	size := int32(len(in))

	return aret, size, size
}

func TransOutBytesReturn(store *wasmtime.Store, instance *wasmtime.Instance, cap int32) (_, _, _ interface{}) {
	// malloc memory space.
	calloc := instance.GetFunc(store, "malloc")
	aret, _ := calloc.Call(store, cap)

	store.InterruptHandle()

	return aret, int32(0), cap
}

func ReadOutBytesReturn(store *wasmtime.Store, instance *wasmtime.Instance, ptr, size int32) []byte {
	return instance.GetExport(store, "memory").Memory().UnsafeData(store)[ptr : ptr+size]
}
