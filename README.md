# go-wasm-go

experiment golang call wasm, and writing wasm in golang

## 测试验证

以fib算法来验证,不管是并发还是单线程速度, wasm机制 + wasmer-go 性能最好,接近源生性能。


### 普通算法

测试并没有多少问题

### 网络发包

最普通的是 http请求

#### tinygo 遇到问题

还没有涉及到调用，仅仅是定义了就遇到问题:

**wazero:**
```go
resolve imports: clock_time_get: not exported in module wasi_snapshot_preview1

=== RUN   Test_wazero_tinygo_Fibonacci
    /data/apps/github.com/wuhuizuo/go-wasm-go/runner/wazero.go:29: resolve imports: clock_time_get: not exported in module wasi_snapshot_preview1
--- FAIL: Test_wazero_tinygo_Fibonacci (0.01s)
```

**wasmer:**
```go
Missing import: `env`.`time.resetTimer`

=== RUN   Test_wasmer_tinygo_Fibonacci
    /data/apps/github.com/wuhuizuo/go-wasm-go/runner/wasmer_tinygo.go:36: Missing import: `env`.`time.resetTimer`
--- FAIL: Test_wasmer_tinygo_Fibonacci (2.52s)
FAIL
FAIL    github.com/wuhuizuo/go-wasm-go/runner   2.561s
```

**wasmedge:**

```go
[2021-12-15 20:08:57.648] [error] instantiation failed: unknown import, Code: 0x62
[2021-12-15 20:08:57.648] [error]     When linking module: "env" , function name: "time.resetTimer"
[2021-12-15 20:08:57.648] [error]     At AST node: import description
[2021-12-15 20:08:57.648] [error]     At AST node: import section
[2021-12-15 20:08:57.648] [error]     At AST node: module
--- FAIL: Test_wasmedge_tinygo_Fibonacci (0.08s)
    wasmedge.go:28: unknown import
```

定义:
```go
//export RequestHTTP
func RequestHTTP(username, password string) {
    provider.RequestHTTP(username, password)
}
```

可能可以参照这个实现进行导入 host 的东西。
https://github.com/mosn/proxy-wasm-go-host/blob/3fb13ba763a662bde51f0f324d465d04d8458449/proxywasm/v2/imports.go#L22

#### go 遇到问题

当相关的包中import 了 `net/http` package后, 会遇到问题:
```go
panic: syscall/js: call of Value.Get on undefined

goroutine 1 [running]:
syscall/js.Value.Get({{}, 0x0, 0x0}, {0x384f8, 0x9})
        /usr/local/go/src/syscall/js/js.go:299 +0xc
syscall.init()
        /usr/local/go/src/syscall/fs_js.go:21 +0x12
FAIL    github.com/wuhuizuo/go-wasm-go/runner   0.714s
```

即时没有使用下面的语句申明也是一样有问题:
```go
js.Global().Set("RequestHTTP", js.FuncOf(wasmutil.Wrap(provider.RequestHTTP)))
```




