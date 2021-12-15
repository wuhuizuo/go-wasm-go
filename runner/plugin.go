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

// newGoPluginFibonacciFn prepare for go plugin `Fibonacci` func.
func newGoPluginFibonacciFn(t testing.TB, soPath, fnName string) func(int32) int32 {
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
func newGoPluginHTTPFn(t testing.TB, soPath, fnName string) func(string, string) {
	f := loadPluginSymbol(t, soPath, fnName)
	switch v := f.(type) {
	case func(string, string):
		return v
	case *func(string, string):
		return *v
	default:
		return nil
	}
}
