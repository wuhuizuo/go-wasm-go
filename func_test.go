package main

import (
	"fmt"
	"testing"

	"github.com/wuhuizuo/go-wasm-go/provider/jsgoja"
)

const (
	goPluginSo = "./provider/plugin/plugin.so"
	wasmTinygo = "./provider/wasm-tinygo/wasm.wasm"
	wasmGo     = "./provider/wasm-go/wasm.wasm"
)

func TestNative_fibonacci(t *testing.T) {
	tests := []struct {
		name string
		in   uint32
		want uint32
	}{
		{name: "20", in: 20, want: 6765},
		{name: "10", in: 10, want: 55},
		{name: "5", in: 5, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.in); got != tt.want {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlugin_fibonacci(t *testing.T) {
	f := newGoPluginFunc(t, goPluginSo, "Fibonacci")

	tests := []struct {
		name string
		in   uint32
		want uint32
	}{
		{name: "20", in: 20, want: 6765},
		{name: "10", in: 10, want: 55},
		{name: "5", in: 5, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJS_fibonacci(t *testing.T) {
	f := jsgoja.NewFibonacci()

	tests := []struct {
		name string
		in   uint32
		want uint32
	}{
		{name: "20", in: 20, want: 6765},
		{name: "10", in: 10, want: 55},
		{name: "5", in: 5, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wazero_Fibonacci(t *testing.T) {
	tests := []struct {
		name string
		in   uint32
		want uint32
	}{
		{name: "20", in: 20, want: 6765},
		{name: "10", in: 10, want: 55},
		{name: "5", in: 5, want: 5},
	}

	t.Run("tinygo", func(t *testing.T) {
		store := newWASMStoreWithWazero(t, wasmTinygo)
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s-%s", t.Name(), tt.name), func(t *testing.T) {
				if got := callWASMFuncWithWazero(t, store, tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("go", func(t *testing.T) {
		t.Skip()
		store := newWASMStoreWithWazero(t, wasmGo)
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s-%s", t.Name(), tt.name), func(t *testing.T) {
				if got := callWASMFuncWithWazero(t, store, tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func Test_wasmer_Fibonacci(t *testing.T) {
	tests := []struct {
		name string
		in   uint32
		want uint32
	}{
		{name: "20", in: 20, want: 6765},
		{name: "10", in: 10, want: 55},
		{name: "5", in: 5, want: 5},
	}

	t.Run("tinygo", func(t *testing.T) {
		fn := getWasmFuncWithWasmer(t, wasmTinygo)
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s-%s", t.Name(), tt.name), func(t *testing.T) {
				if got := callWASMFuncWithWasmer(t, fn, tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("go", func(t *testing.T) {
		t.Skip()
		fn := getWasmFuncWithWasmer(t, wasmGo)
		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s-%s", t.Name(), tt.name), func(t *testing.T) {
				if got := callWASMFuncWithWasmer(t, fn, tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func Test_wasmedge_Fibonacci(t *testing.T) {
	tests := []struct {
		name string
		in   uint32
		want uint32
	}{
		{name: "20", in: 20, want: 6765},
		{name: "10", in: 10, want: 55},
		{name: "5", in: 5, want: 5},
	}

	t.Run("tinygo", func(t *testing.T) {
		vm, conf := getWasmedgeInstance(t, wasmTinygo)
		defer vm.Release()
		defer conf.Release()

		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s-%s", t.Name(), tt.name), func(t *testing.T) {
				if got := callWASMFuncWithWasmedge(t, vm, tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("go", func(t *testing.T) {
		t.Skip()
		vm, conf := getWasmedgeInstance(t, wasmGo)
		defer vm.Release()
		defer conf.Release()

		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s-%s", t.Name(), tt.name), func(t *testing.T) {
				if got := callWASMFuncWithWasmedge(t, vm, tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}
