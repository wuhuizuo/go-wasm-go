package runner

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmer"
)

func Test_wasmer_tinygo(t *testing.T) {
	t.Run("algorithm", func(t *testing.T) {
		t.Skip(`
			fatal error: unexpected signal during runtime execution
			[signal SIGSEGV: segmentation violation code=0x1 addr=0x1560ff6c00 pc=0x7f745455115e]

			runtime stack:
			runtime.throw({0x1737f15, 0x0})
				/usr/local/go/src/runtime/panic.go:1198 +0x71
			runtime.sigpanic()
				/usr/local/go/src/runtime/signal_unix.go:719 +0x396
		`)

		fn := wasmer.GetWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasi), fibFuncName)

		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := wasmer.CallWASMFuncWithWasmer(t, fn, tt.in); got.(int32) != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// t.Run("http request", func(t *testing.T) {
	// 	fn := getWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmTinygo), httpReqFuncName)
	// 	if got := callWASMFuncWithWasmer(t, fn); got != nil {
	// 		t.Errorf("fn() = %v, want %v", got, nil)
	// 	}
	// })
}

func Test_wasmer_go(t *testing.T) {
	t.Run("algorithm", func(t *testing.T) {
		fn := wasmer.GetGoWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmGo), fibFuncName)

		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := wasmer.CallGoWASMFuncWithWasmer(t, fn, []interface{}{tt.in}).(float64); int32(got) != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		fn := wasmer.GetGoWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmGo), httpReqFuncName)
		wasmer.CallGoWASMFuncWithWasmer(t, fn, nil)
	})

	t.Run("file io", func(t *testing.T) {
		fn := wasmer.GetGoWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmGo), ioFunName)
		got := wasmer.CallGoWASMFuncWithWasmer(t, fn, nil)
		assert.Equal(t, got, 0)
	})

	t.Run("multi threads", func(t *testing.T) {
		fn := wasmer.GetGoWasmFuncWithWasmer(t, filepath.Join(selfDir(t), "..", wasmGo), multiThreadsFuncName)
		wasmer.CallGoWASMFuncWithWasmer(t, fn, []interface{}{4})
	})
}
