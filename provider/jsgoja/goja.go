package jsgoja

import (
	_ "embed"

	"github.com/dop251/goja"
)

//go:embed goja.js
var jsScript string

func NewFibonacci() func(uint32) uint32 {
	vm := goja.New()
	_, err := vm.RunString(jsScript)
	if err != nil {
		panic(err)
	}

	var fn func(uint32) uint32
	if err := vm.ExportTo(vm.Get("fibonacci"), &fn); err != nil {
		panic(err)
	}

	return fn
}
