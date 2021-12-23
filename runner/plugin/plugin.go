package plugin

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

// NewGoPluginAlgFn prepare for go plugin `Fibonacci` func.
func NewGoPluginAlgFn(t testing.TB, soPath, fnName string) func(int32) int32 {
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

// NewGoPluginIOFn prepare for go plugin http test func.
func NewGoPluginIOFn(t testing.TB, soPath, fnName string) func() {
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

// NewGoPluginIOErrFn prepare for go plugin http test func.
func NewGoPluginIOErrFn(t testing.TB, soPath, fnName string) func() error {
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

// NewGoPluginMultiThreads prepare for go plugin multi threads test func.
func NewGoPluginMultiThreads(t testing.TB, soPath, fnName string) func(int32) {
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
