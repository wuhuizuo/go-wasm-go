package runner

import (
	"path/filepath"
	"testing"

	"github.com/wuhuizuo/go-wasm-go/provider/jsgoja"
	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

func TestNative_algorithm(t *testing.T) {
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
		native.RequestHTTP()
	})

	t.Run("file io", func(t *testing.T) {
		if err := native.FileIO(); err != nil {
			t.Error(err)
		}
	})

	t.Run("multi threads", func(t *testing.T) {
		native.MultiThreads(4)
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
				if got := callWASMFuncWithWazero(t, store, tt.in); got != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		// TODO:
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
		// TODO:
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
