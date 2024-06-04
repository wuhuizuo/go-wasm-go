package runner

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmtime"
)

func Test_wasmtime(t *testing.T) {
	for _, f := range []string{wasmTinygo, wasmGo} {
		t.Run(fmt.Sprintf("test %s", f), func(t *testing.T) {
			store, instance := wasmtime.GetWasmFuncWithWasmtime(t, filepath.Join(selfDir(t), "..", f))

			t.Run("algorithm", func(t *testing.T) {
				for _, tt := range fibTests {
					t.Run(tt.name, func(t *testing.T) {
						got, err := wasmtime.CallWasmFunc(t, store, instance, fibFuncName, tt.in)

						assert.NoError(t, err)
						assert.IsType(t, tt.want, got)
						assert.Equal(t, tt.want, got)
					})
				}
			})

			t.Run("bytes test", func(t *testing.T) {
				inPtr, inSize, inCap := wasmtime.TransInBytesParam(store, instance, []byte("hello"))
				got, err := wasmtime.CallWasmFunc(t, store, instance, byteInOutFuncName, inPtr, inSize, inCap)
				t.Log(got, err)
				assert.NoError(t, err)

				outSize, err := wasmtime.CallWasmFunc(t, store, instance, byteInOutLenFuncName)
				t.Log(outSize, err)
				assert.NoError(t, err)

				out := wasmtime.ReadOutBytesReturn(store, instance, got.(int32), outSize.(int32))
				assert.Equal(t, []byte("hello---"), out)
			})
		})
	}
}
