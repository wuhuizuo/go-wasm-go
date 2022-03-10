package runner

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmtime"
)

func Test_wasmtime_tinygo(t *testing.T) {
	store, instance := wasmtime.GetRuntimes(t, filepath.Join(selfDir(t), "..", wasi), nil)

	t.Run("algorithm", func(t *testing.T) {
		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := wasmtime.CallFunc(t, store, instance, fibFuncName, tt.in)

				assert.NoError(t, err)
				assert.IsType(t, tt.want, got)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("bytes test", func(t *testing.T) {
		inPtr, inSize, inCap := wasmtime.TransInBytesParam(store, instance, []byte("hello"))
		got, err := wasmtime.CallFunc(t, store, instance, byteInOutFuncName, inPtr, inSize, inCap)
		t.Log(got, err)
		assert.NoError(t, err)

		outPtr, outLen := int32(got.(int64)>>32), int32(got.(int64))
		out := wasmtime.ReadOutBytesReturn(store, instance, outPtr, outLen)
		assert.Equal(t, []byte("hello---"), out)
	})

	t.Run("string test", func(t *testing.T) {
		inPtr, inSize := wasmtime.TransInStringParam(store, instance, "hello")
		got, err := wasmtime.CallFunc(t, store, instance, stringInOutFuncName, inPtr, inSize)
		t.Log(got, err)
		assert.NoError(t, err)

		outPtr, outLen := int32(got.(int64)>>32), int32(got.(int64))
		out := wasmtime.ReadOutStringReturn(store, instance, outPtr, outLen)
		assert.Equal(t, "hello---", out)
	})
}

func Test_wasmtime_tinygo_moduleLinking(t *testing.T) {
	deps := map[string]string{
		"standalone": filepath.Join(selfDir(t), "..", wasi),
	}
	store, instance := wasmtime.GetRuntimes(t, filepath.Join(selfDir(t), "..", wasiParent), deps)

	t.Run("algorithm", func(t *testing.T) {
		for _, tt := range fibTests {
			t.Run(tt.name, func(t *testing.T) {
				got, err := wasmtime.CallFunc(t, store, instance, fibFuncName, tt.in)

				assert.NoError(t, err)
				assert.IsType(t, tt.want, got)
				assert.Equal(t, tt.want, got)
			})
		}
	})

	t.Run("bytes test", func(t *testing.T) {
		t.Skip("TODO:read/write memory cross modules was not supported")
		inPtr, inSize, inCap := wasmtime.TransInBytesParam(store, instance, []byte("hello"))
		got, err := wasmtime.CallFunc(t, store, instance, byteInOutFuncName, inPtr, inSize, inCap)
		t.Log(got, err)
		assert.NoError(t, err)

		outPtr, outLen := int32(got.(int64)>>32), int32(got.(int64))
		out := wasmtime.ReadOutBytesReturn(store, instance, outPtr, outLen)
		assert.Equal(t, []byte("hello---"), out)
	})

	t.Run("string test", func(t *testing.T) {
		t.Skip("TODO:read/write memory cross modules was not supported")
		inPtr, inSize := wasmtime.TransInStringParam(store, instance, "hello")
		got, err := wasmtime.CallFunc(t, store, instance, stringInOutFuncName, inPtr, inSize)
		t.Log(got, err)
		assert.NoError(t, err)

		outPtr, outLen := int32(got.(int64)>>32), int32(got.(int64))
		out := wasmtime.ReadOutStringReturn(store, instance, outPtr, outLen)
		assert.Equal(t, "hello---", out)
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

func checkerr(err error) {
	if err != nil {
		panic(err)
	}
}
