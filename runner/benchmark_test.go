package runner

import (
	"fmt"
	"path/filepath"
	"runtime/debug"
	"testing"

	"github.com/wuhuizuo/go-wasm-go/provider/jsgoja"
	"github.com/wuhuizuo/go-wasm-go/provider/native"
	"github.com/wuhuizuo/go-wasm-go/runner/plugin"
	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmedge"
	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmer"
	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmtime"
	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wazero"
)

func Benchmark_fibonacci_single_10(b *testing.B) {
	benchmark_fibonacci_single(b, 10)
}

func Benchmark_fibonacci_single_20(b *testing.B) {
	benchmark_fibonacci_single(b, 20)
}

func Benchmark_fibonacci_single_30(b *testing.B) {
	benchmark_fibonacci_single(b, 30)
}

func Benchmark_fibonacci_single_40(b *testing.B) {
	benchmark_fibonacci_single(b, 40)
}

func Benchmark_fibonacci_paralle_10(b *testing.B) {
	benchmark_fibonacci_paralle(b, 10)
}

func Benchmark_fibonacci_paralle_20(b *testing.B) {
	benchmark_fibonacci_paralle(b, 20)
}

func Benchmark_fibonacci_paralle_30(b *testing.B) {
	benchmark_fibonacci_paralle(b, 30)
}

func Benchmark_fibonacci_paralle_40(b *testing.B) {
	benchmark_fibonacci_paralle(b, 40)
}

func benchmark_fibonacci_paralle(b *testing.B, fbIn int32) {
	// go gc 会导致 wasmer 异常,需要手动线停止GC.
	debug.SetGCPercent(-1)
	b.ResetTimer()

	b.Run(fmt.Sprintf("native - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				native.Fibonacci(fbIn)
			}
		})
	})

	b.Run(fmt.Sprintf("plugin - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			soFn := plugin.NewGoPluginAlgFn(b, filepath.Join(selfDir(b), "..", goPluginSo), fibFuncName)
			for pb.Next() {
				soFn(fbIn)
			}
		})
	})

	b.Run(fmt.Sprintf("wasm-wasmtime - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			store, instance := wasmtime.GetWasmFuncWithWasmtime(b, filepath.Join(selfDir(b), "..", wasmTinygo))

			for pb.Next() {
				wasmtime.CallWasmFunc(b, store, instance, fibFuncName)
			}
		})

	})

	b.Run(fmt.Sprintf("wasm-wazero - fb(%d)", fbIn), func(b *testing.B) {
		mod, closer := wazero.NewWASMStoreWithWazero(b, filepath.Join(selfDir(b), "..", wasmTinygo))
		defer closer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				wazero.CallWASMFuncWithWazero(b, mod, fibFuncName, uint64(fbIn))
			}
		})
	})

	b.Run(fmt.Sprintf("wasm-wasmer - fb(%d)", fbIn), func(b *testing.B) {
		b.Skip()
		b.RunParallel(func(pb *testing.PB) {
			wasmFn := wasmer.GetWasmFuncWithWasmer(b, filepath.Join(selfDir(b), "..", wasmTinygo), fibFuncName)
			for pb.Next() {
				wasmer.CallWASMFuncWithWasmer(b, wasmFn, fbIn)
			}
		})
	})

	b.Run(fmt.Sprintf("wasm-wasmedge - fb(%d)", fbIn), func(b *testing.B) {
		b.Skip("unexpected signal during runtime execution")
		b.RunParallel(func(pb *testing.PB) {
			vm, conf := wasmedge.GetWasmedgeInstance(b, filepath.Join(selfDir(b), "..", wasmTinygo))
			defer vm.Release()
			defer conf.Release()
			for pb.Next() {
				wasmedge.CallWASMFuncWithWasmedgeReturnInt32(b, vm, fibFuncName, fbIn)
			}
		})
	})

	b.Run(fmt.Sprintf("js - fb(%d)", fbIn), func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			f := jsgoja.NewFibonacci()
			for pb.Next() {
				f(fbIn)
			}
		})
	})
}

func benchmark_fibonacci_single(b *testing.B, fbIn int32) {
	b.Run(fmt.Sprintf("native - fb(%d)", fbIn), func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			native.Fibonacci(fbIn)
		}
	})

	b.Run(fmt.Sprintf("plugin - fb(%d)", fbIn), func(b *testing.B) {
		soFn := plugin.NewGoPluginAlgFn(b, filepath.Join(selfDir(b), "..", goPluginSo), fibFuncName)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			soFn(fbIn)
		}
	})

	b.Run(fmt.Sprintf("wasm-wasmtime - fb(%d)", fbIn), func(b *testing.B) {
		store, instance := wasmtime.GetWasmFuncWithWasmtime(b, filepath.Join(selfDir(b), "..", wasmTinygo))
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			wasmtime.CallWasmFunc(b, store, instance, fibFuncName, fibFuncName, fbIn)
		}
	})

	b.Run(fmt.Sprintf("wasm-wazero - fb(%d)", fbIn), func(b *testing.B) {
		mod, closer := wazero.NewWASMStoreWithWazero(b, filepath.Join(selfDir(b), "..", wasmTinygo))
		defer closer()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			wazero.CallWASMFuncWithWazero(b, mod, fibFuncName, uint64(fbIn))
		}
	})

	b.Run(fmt.Sprintf("wasm-wasmer - fb(%d)", fbIn), func(b *testing.B) {
		b.Skip("unexpected signal during runtime execution")
		wasmFn := wasmer.GetWasmFuncWithWasmer(b, filepath.Join(selfDir(b), "..", wasmTinygo), fibFuncName)
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			wasmer.CallWASMFuncWithWasmer(b, wasmFn, fbIn)
		}
	})
	b.Run(fmt.Sprintf("wasm-wasmedge - fb(%d)", fbIn), func(b *testing.B) {
		vm, conf := wasmedge.GetWasmedgeInstance(b, filepath.Join(selfDir(b), "..", wasmTinygo))
		defer vm.Release()
		defer conf.Release()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			wasmedge.CallWASMFuncWithWasmedgeReturnInt32(b, vm, fibFuncName, fbIn)
		}
	})

	b.Run(fmt.Sprintf("js - fb(%d)", fbIn), func(b *testing.B) {
		f := jsgoja.NewFibonacci()
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			f(fbIn)
		}
	})
}
