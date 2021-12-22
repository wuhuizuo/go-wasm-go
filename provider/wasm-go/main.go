package main

import (
	"reflect"
	"syscall/js"

	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

func main() {
	js.Global().Set("Fibonacci", js.FuncOf(Wrap(native.Fibonacci)))
	js.Global().Set("RequestHTTP", js.FuncOf(Wrap(native.RequestHTTP)))
	js.Global().Set("FileIO", js.FuncOf(Wrap(native.FileIO)))
	js.Global().Set("MultiThreads", js.FuncOf(Wrap(native.MultiThreads)))

	select {}
}

// Wrap make function for friendly to js/wasm.
func Wrap(f interface{}) func(js.Value, []js.Value) interface{} {
	return func(this js.Value, args []js.Value) interface{} {
		rf := reflect.ValueOf(f)
		rt := rf.Type()
		rargs := []reflect.Value{}
		for i := 0; i < rt.NumIn(); i++ {
			var arg reflect.Value
			switch args[i].Type() {
			case js.TypeUndefined:
				arg = reflect.Zero(reflect.TypeOf(nil)).Convert(rt.In(i))
			case js.TypeNull:
				arg = reflect.Zero(reflect.TypeOf(nil)).Convert(rt.In(i))
			case js.TypeBoolean:
				arg = reflect.ValueOf(args[i].Bool()).Convert(rt.In(i))
			case js.TypeNumber:
				arg = reflect.ValueOf(args[i].Float()).Convert(rt.In(i))
			case js.TypeString:
				arg = reflect.ValueOf(args[i].String()).Convert(rt.In(i))
			case js.TypeSymbol:
				arg = reflect.ValueOf(args[i].String()).Convert(rt.In(i))
			case js.TypeObject:
				arg = reflect.ValueOf(args[i].JSValue()).Convert(rt.In(i))
			case js.TypeFunction:
				arg = reflect.ValueOf(args[i].JSValue()).Convert(rt.In(i))
			default:
				arg = reflect.ValueOf(args[i].JSValue()).Convert(rt.In(i))
			}
			rargs = append(rargs, arg)
		}
		ret := rf.Call(rargs)
		if len(ret) > 0 {
			return ret[0].Interface()
		}
		return nil
	}
}
