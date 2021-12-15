package runner

import (
	"path/filepath"
	"testing"

	"github.com/wuhuizuo/go-wasm-go/provider"
	"github.com/wuhuizuo/go-wasm-go/provider/jsgoja"
)

func TestNative_Fibonacci(t *testing.T) {
	for _, tt := range fbTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := provider.Fibonacci(tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNative_HTTPBasicAuth(t *testing.T) {
	provider.HTTPBasicAuth("xxx", "yyy")
}

func TestPlugin_Fibonacci(t *testing.T) {
	f := newGoPluginFibonacciFn(t, filepath.Join(selfDir(t), "..", goPluginSo), fibFuncName)

	for _, tt := range fbTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlugin_HTTPBasicAuth(t *testing.T) {
	f := newGoPluginHTTPFn(t, filepath.Join(selfDir(t), "..", goPluginSo), httpReqFuncName)
	f("xxx", "yyy")
}

func TestJS_Fibonacci(t *testing.T) {
	f := jsgoja.NewFibonacci()

	for _, tt := range fbTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f(tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJS_HTTPBasicAuth(t *testing.T) {
	t.Skip("不支持http请求")

	f := jsgoja.NewHTTPBasicAuth()
	f("xxx", "yyy")
}

func Test_wazero_tinygo_Fibonacci(t *testing.T) {
	store := newWASMStoreWithWazero(t, filepath.Join(selfDir(t), "..", wasmTinygo))

	for _, tt := range fbTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callWASMFuncWithWazero(t, store, tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wasmer_tinygo_Fibonacci(t *testing.T) {
	fn := getWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmTinygo))

	for _, tt := range fbTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callWASMFuncWithWasmer(t, fn, tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wasmer_go_Fibonacci(t *testing.T) {
	fn := getGoWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmGo))

	for _, tt := range fbTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callGoWASMFuncWithWasmer(t, fn, tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wasmedge_tinygo_Fibonacci(t *testing.T) {
	vm, conf := getWasmedgeInstance(t, filepath.Join(selfDir(t), "..", wasmTinygo))
	defer vm.Release()
	defer conf.Release()

	for _, tt := range fbTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := callWASMFuncWithWasmedgeReturnInt32(t, vm, fibFuncName, tt.in); got != tt.want {
				t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func Test_wasmedge_tinygo_HTTP(t *testing.T) {
// 	vm, conf := getWasmedgeInstance(t, filepath.Join(selfDir(t), "..", wasmTinygo))
// 	defer vm.Release()
// 	defer conf.Release()

// 	callWASMFuncWithWasmedgeReturnVoid(t, vm, httpReqFuncName, "xxx", "yyy")
// }
