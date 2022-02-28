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

	t.Run("arg types", func(t *testing.T) {
		t.Skip("FIXME: how to passing byte slice args to wasm func.")
		store, fn := wasmtime.GetWasmFuncWithWasmtime(t, filepath.Join(selfDir(t), "..", wasmTinygo), typeFuncName)

		got, err := fn.Call(store, 0, 0, 0, 0)
		t.Log(got)
		assert.NoError(t, err)
		assert.True(t, false)
	})
}

func Test_wasmtime_go(t *testing.T) {
	t.Skip()
	t.Run("algorithm", func(t *testing.T) {
		store, fn := wasmtime.GetGoWasmFuncWithWasmtime(t, filepath.Join(selfDir(t), "..", wasmGo), fibFuncName)

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
