package runner

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wazero"
)

func Test_wazero_tinygo(t *testing.T) {
	mod, closer := wazero.NewWASMStoreWithWazero(t, filepath.Join(selfDir(t), "..", wasmTinygo))
	defer closer()

	t.Run("algorithm", func(t *testing.T) {
		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				if got := wazero.CallWASMFuncWithWazero(t, mod, fibFuncName, uint64(tt.in)); int32(got[0]) != tt.want {
					t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	t.Run("http request", func(t *testing.T) {
		got := wazero.CallWASMFuncWithWazero(t, mod, httpReqFuncName)
		assert.Len(t, got, 1)
	})

	t.Run("file io", func(t *testing.T) {
		got := wazero.CallWASMFuncWithWazero(t, mod, ioFunName)
		assert.Len(t, got, 1)
	})

	t.Run("multi threads", func(t *testing.T) {
		got := wazero.CallWASMFuncWithWazero(t, mod, multiThreadsFuncName, 4)
		assert.Len(t, got, 1)
	})
}

// NOTE: GOOS=js works in wazero, but dynamic functions are not yet supported.
// See https://github.com/tetratelabs/wazero/tree/main/imports/go/example
