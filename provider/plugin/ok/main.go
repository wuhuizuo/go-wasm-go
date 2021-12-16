package main

import (
	"github.com/wuhuizuo/go-wasm-go/provider/native"
)

var (
	Fibonacci    = native.Fibonacci
	RequestHTTP  = native.RequestHTTP
	MultiThreads = native.MultiThreads
	FileIO       = native.FileIO
)
