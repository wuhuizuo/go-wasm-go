package runner

import (
	"plugin"
	"testing"
)

func loadPluginSymbol(t testing.TB, soPath, fnName string) plugin.Symbol {
	p, err := plugin.Open(soPath)
	if err != nil {
		t.Fatal(err)
	}

	s, err := p.Lookup(fnName)
	if err != nil {
		t.Fatal(err)
	}

	return s
}

// newGoPluginAlgFn prepare for go plugin `Fibonacci` func.
func newGoPluginAlgFn(t testing.TB, soPath, fnName string) func(int32) int32 {
	f := loadPluginSymbol(t, soPath, fnName)
	switch v := f.(type) {
	case func(int32) int32:
		return v
	case *func(int32) int32:
		return *v
	default:
		return nil
	}
}

// newGoPluginFibonacciFn prepare for go plugin http test func.
func newGoPluginIOFn(t testing.TB, soPath, fnName string) func() {
	f := loadPluginSymbol(t, soPath, fnName)
	switch v := f.(type) {
	case func():
		return v
	case *func():
		return *v
	default:
		return nil
	}
}

// newGoPluginFibonacciFn prepare for go plugin http test func.
func newGoPluginIOErrFn(t testing.TB, soPath, fnName string) func() error {
	f := loadPluginSymbol(t, soPath, fnName)
	switch v := f.(type) {
	case func() error:
		return v
	case *func() error:
		return *v
	default:
		return nil
	}
}

func newGoPluginMultiThreads(t testing.TB, soPath, fnName string) func(int32) {
	f := loadPluginSymbol(t, soPath, fnName)
	switch v := f.(type) {
	case func(int32):
		return v
	case *func(int32):
		return *v
	default:
		return nil
	}
}
