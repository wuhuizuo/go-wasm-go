package main

import (
	"syscall/js"
)

func main() {
	js.Global().Set("Fibonacci", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		return Fibonacci(uint32(args[0]))
	}))
	select {}
}

func Fibonacci(in uint32) uint32 {
	if in <= 1 {
		return in
	}
	return Fibonacci(in-1) + Fibonacci(in-2)
}
