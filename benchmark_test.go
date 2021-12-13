package main

import "testing"

func Benchmark_fibonacci_single(b *testing.B) {
	// 太大了,会卡死.
	fbIn := uint32(30)
	b.N = 1000

	b.Run("go native", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fibonacci(fbIn)
		}
	})

	soFn := newGoPluginFunc(b, goPluginSo, wasmFuncName)
	b.ResetTimer()
	b.Run("go plugin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			soFn(fbIn)
		}
	})

	store := newWASMStoreWithWazero(b, wasmTinygo)
	b.ResetTimer()
	b.Run("wasm - tinygo - wazero", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			callWASMFuncWithWazero(b, store, fbIn)
		}
	})

	// wasmFn := getWasmFuncWithWasmer(b, wasmTinygo)
	// b.ResetTimer()
	// b.Run("wasm - tinygo - wasmer", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		b.Log(i)
	// 		callWASMFuncWithWasmer(b, wasmFn, fbIn)
	// 	}
	// })
}

func Benchmark_fibonacci_paralle(b *testing.B) {
	// 太大了,会卡死.
	fbIn := uint32(30)
	b.N = 1
	par := 4

	b.Run("go native", func(b *testing.B) {
		b.SetParallelism(par)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				fibonacci(fbIn)
			}
		})
	})

	soFn := newGoPluginFunc(b, goPluginSo, wasmFuncName)
	b.ResetTimer()
	b.Run("go plugin", func(b *testing.B) {
		b.SetParallelism(par)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				soFn(fbIn)
			}
		})
	})

	b.ResetTimer()
	b.Run("wasm - tinygo - wazero", func(b *testing.B) {
		b.SetParallelism(par)
		b.RunParallel(func(pb *testing.PB) {
			// 必须在线程里面加载, 不能在线程外加载，然后并发.
			store := newWASMStoreWithWazero(b, wasmTinygo)
			for pb.Next() {
				callWASMFuncWithWazero(b, store, fbIn)
			}
		})
	})

	// wasmFn := getWasmFuncWithWasmer(b, wasmTinygo)
	// b.ResetTimer()
	// b.Run("wasm - tinygo - wasmer", func(b *testing.B) {
	// 	for i := 0; i < b.N; i++ {
	// 		b.Log(i)
	// 		callWASMFuncWithWasmer(b, wasmFn, fbIn)
	// 	}
	// })
}
