package runner

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wuhuizuo/go-wasm-go/provider/jsgoja"
	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

func TestNative(t *testing.T) {
	t.Run("algorithm", func(t *testing.T) {
		for _, tt := range fbTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := native.Fibonacci(tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		if got := native.RequestHTTP(); got != 0 {
			t.Error(got)
		}
	})

	t.Run("file io", func(t *testing.T) {
		if got := native.FileIO(); got != 0 {
			t.Error(got)
		}
	})

	t.Run("multi threads", func(t *testing.T) {
		if got := native.MultiThreads(4); got != 0 {
			t.Error(got)
		}
	})
}

func TestJS(t *testing.T) {
	t.Run("algorithm", func(t *testing.T) {
		f := jsgoja.NewFibonacci()

		for _, tt := range fbTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := f(tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		t.Skip("不支持http请求")

		f := jsgoja.NewRequestHTTP()
		f("xxx", "yyy")
	})
}

func Test_wazero_tinygo(t *testing.T) {
	store := newWASMStoreWithWazero(t, filepath.Join(selfDir(t), "..", wasmTinygo))

	t.Run("algorithm", func(t *testing.T) {
		for _, tt := range fbTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := callWASMFuncWithWazero(t, store, fibFuncName, uint64(tt.in)); int32(got[0]) != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		got := callWASMFuncWithWazero(t, store, httpReqFuncName, 1)
		assert.Len(t, got, 1)
	})

	t.Run("file io", func(t *testing.T) {
		got := callWASMFuncWithWazero(t, store, ioFunName, 1)
		assert.Len(t, got, 1)
	})

	t.Run("multi threads", func(t *testing.T) {
		got := callWASMFuncWithWazero(t, store, multiThreadsFuncName, 4)
		assert.Len(t, got, 1)
	})
}

func Test_wasmer_tinygo(t *testing.T) {
	t.Run("algorithm", func(t *testing.T) {
		fn := getWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmTinygo), fibFuncName)

		for _, tt := range fbTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := callWASMFuncWithWasmer(t, fn, tt.in); got.(int32) != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		fn := getWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmTinygo), httpReqFuncName)
		if got := callWASMFuncWithWasmer(t, fn); got != nil {
			t.Errorf("fn() = %v, want %v", got, nil)
		}
	})
}

func Test_wasmer_go(t *testing.T) {
	t.Run("algorithm", func(t *testing.T) {
		fn := getGoWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmGo), fibFuncName)

		for _, tt := range fbTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := callGoWASMFuncWithWasmer(t, fn, []interface{}{tt.in}).(float64); int32(got) != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		fn := getGoWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmGo), httpReqFuncName)
		callGoWASMFuncWithWasmer(t, fn, nil)
	})

	t.Run("file io", func(t *testing.T) {
		fn := getGoWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmGo), ioFunName)
		got := callGoWASMFuncWithWasmer(t, fn, nil)
		assert.Equal(t, got, 0)
	})

	t.Run("multi threads", func(t *testing.T) {
		fn := getGoWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmGo), multiThreadsFuncName)
		callGoWASMFuncWithWasmer(t, fn, []interface{}{4})
	})
}

func Test_wasmtime_tinygo(t *testing.T) {
	t.Skip("expected 1 imports, found 0")

	t.Run("algorithm", func(t *testing.T) {
		store, fn := getWasmFuncWithWasmtime(t, filepath.Join(selfDir(t), "..", wasmTinygo), fibFuncName)

		for _, tt := range fbTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := fn.Call(store, tt.in)
				if err != nil {
					t.Error(err)
				}

				if int32(got.(float64)) != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

}

func Test_wasmedge_tinygo(t *testing.T) {
	vm, conf := getWasmedgeInstance(t, filepath.Join(selfDir(t), "..", wasmTinygo))
	defer vm.Release()
	defer conf.Release()

	t.Run("algorithm", func(t *testing.T) {
		for _, tt := range fbTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := callWASMFuncWithWasmedgeReturnInt32(t, vm, fibFuncName, tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		callWASMFuncWithWasmedgeReturnVoid(t, vm, httpReqFuncName)
	})
}
