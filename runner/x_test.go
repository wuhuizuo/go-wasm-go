package runner

import (
	"path/filepath"
	"runtime"
	"testing"
)

const (
	wasmTinygo = "provider/wasm-tinygo/wasm.wasm"
	wasmGo     = "provider/wasm-go/wasm.wasm"
)

const (
	fibFuncName          = "Fibonacci"
	httpReqFuncName      = "RequestHTTP"
	ioFunName            = "FileIO"
	multiThreadsFuncName = "MultiThreads"
	byteInOutFuncName    = "BytesTest"
	byteInOutLenFuncName = "BytesTestLen"
)

var fibTests = []fibTestItem{
	{name: "5", in: 5, want: 5},
	{name: "10", in: 10, want: 55},
	{name: "20", in: 20, want: 6765},
	{name: "30", in: 30, want: 832040},
	// {name: "40", in: 40, want: 102334155},
}

type fibTestItem struct {
	name string
	in   int32
	want int32
}

// selfDir get current test file dir path
func selfDir(t testing.TB) string {
	t.Helper()

	// nolint: dogsled
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}
