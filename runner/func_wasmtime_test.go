package runner

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmtime"
)

func Test_wasmtime_tinygo(t *testing.T) {
	t.Run("algorithm", func(t *testing.T) {
		store, fn := wasmtime.GetWasmFuncWithWasmtime(t, filepath.Join(selfDir(t), "..", wasmTinygo), fibFuncName)

		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := fn.Call(store, tt.in)

				assert.NoError(t, err)
				assert.IsType(t, tt.want, got)
				assert.Equal(t, tt.want, got)
			})
		}
	})
}
