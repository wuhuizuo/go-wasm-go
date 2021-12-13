package main

import (
	"plugin"
	"testing"
)

// newGoPluginFunc prepare for go plugin test func.
func newGoPluginFunc(t testing.TB, soPath, fnName string) func(uint32) uint32 {
	p, err := plugin.Open(soPath)
	if err != nil {
		t.Fatal(err)
	}

	f, err := p.Lookup(fnName)
	if err != nil {
		t.Fatal(err)
	}

	return f.(func(uint32) uint32)
}
