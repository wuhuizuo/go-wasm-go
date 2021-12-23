package runner

import (
	"path/filepath"
	"testing"

	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmtime"
)

func Test_wasmtime_tinygo(t *testing.T) {
	t.Skip("expected 1 imports, found 0")

	t.Run("algorithm", func(t *testing.T) {
		store, fn := wasmtime.GetWasmFuncWithWasmtime(t, filepath.Join(selfDir(t), "..", wasmTinygo), fibFuncName)

		for _, tt := range fibTests {
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
