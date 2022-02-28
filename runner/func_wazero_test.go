package runner

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wazero"
)

func Test_wazero_tinygo(t *testing.T) {
	store := wazero.NewWASMStoreWithWazero(t, filepath.Join(selfDir(t), "..", wasmTinygo))

	t.Run("algorithm", func(t *testing.T) {
		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := wazero.CallWASMFuncWithWazero(t, store, fibFuncName, uint64(tt.in)); int32(got[0]) != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		got := wazero.CallWASMFuncWithWazero(t, store, httpReqFuncName)
		assert.Len(t, got, 1)
	})

	t.Run("file io", func(t *testing.T) {
		got := wazero.CallWASMFuncWithWazero(t, store, ioFunName)
		assert.Len(t, got, 1)
	})

	t.Run("multi threads", func(t *testing.T) {
		got := wazero.CallWASMFuncWithWazero(t, store, multiThreadsFuncName, 4)
		assert.Len(t, got, 1)
	})
}

func Test_wazero_go(t *testing.T) {
	t.Skip("not found func")
	store := wazero.NewGoWASMStoreWithWazero(t, filepath.Join(selfDir(t), "..", wasmGo))

	t.Run("algorithm", func(t *testing.T) {
		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := wazero.CallGoWASMFuncWithWazero(t, store, fibFuncName, uint64(tt.in)); int32(got[0]) != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		got := wazero.CallGoWASMFuncWithWazero(t, store, httpReqFuncName)
		assert.Len(t, got, 1)
	})

	t.Run("file io", func(t *testing.T) {
		got := wazero.CallGoWASMFuncWithWazero(t, store, ioFunName)
		assert.Len(t, got, 1)
	})

	t.Run("multi threads", func(t *testing.T) {
		got := wazero.CallGoWASMFuncWithWazero(t, store, multiThreadsFuncName, 4)
		assert.Len(t, got, 1)
	})
}
