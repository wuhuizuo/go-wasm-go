package runner

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/wuhuizuo/go-wasm-go/runner/wasm/wasmedge"
)

func Test_wasmedge(t *testing.T) {
	for _, f := range []string{wasmTinygo, wasmGo} {
		t.Run(fmt.Sprintf("test %s", f), func(t *testing.T) {
			vm, conf := wasmedge.GetWasmedgeInstance(t, filepath.Join(selfDir(t), "..", f))
			defer vm.Release()
			defer conf.Release()

			t.Run("algorithm", func(t *testing.T) {
				for _, tt := range fibTests {
					t.Run(tt.name, func(t *testing.T) {
						if got := wasmedge.CallWASMFuncWithWasmedgeReturnInt32(t, vm, fibFuncName, tt.in); got != tt.want {
							t.Errorf("Fibonacci() = %v, want %v", got, tt.want)
						}
					})
				}
			})

			t.Run("http request", func(t *testing.T) {
				wasmedge.CallWASMFuncWithWasmedgeReturnVoid(t, vm, httpReqFuncName)
			})
		})
	}
}
