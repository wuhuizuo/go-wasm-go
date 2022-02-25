wasmedge_version := "0.9.2"
tinygo_ver := "0.22.0"
 
build: build_wasm_tinygo build_wasm_go build_plugin-all

build_wasm_tinygo:
    #!/usr/bin/env sh
    cd provider/wasm-tinygo &&
    tinygo build -target=wasi -o wasm.wasm

build_wasm_go:
    #!/usr/bin/env sh
    cd provider/wasm-go
    GOOS=js GOARCH=wasm go build -o wasm.wasm

build_plugin-all: 
    just build_plugin plugin/ok
    just build_plugin plugin/third
    just build_plugin plugin/third_diff_mod_ver
    just build_plugin_ver 1.16.14
    just build_plugin_ver 1.17.7

build_plugin TARGET:
    #!/usr/bin/env sh
    cd provider/{{TARGET}}
    go build -buildmode=plugin -o plugin.so

build_plugin_ver VER:
    docker run --rm -v $(pwd):/ws -w /ws/provider/plugin/ok golang:{{VER}} \
       go build -buildmode=plugin -o plugin-{{VER}}.so

install_tools: install_wasmedge install_tinygo

install_wasmedge:
    wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | sudo bash -s -- -v {{wasmedge_version}} -e all -p /usr/local/ -r yes

install_tinygo:    
    wget -qO- https://github.com/tinygo-org/tinygo/releases/download/v{{tinygo_ver}}/tinygo{{tinygo_ver}}.linux-amd64.tar.gz | sudo tar -zxf - -C /usr/local/
    sudo ln -s /usr/local/tinygo/bin/tinygo /usr/local/bin/tinygo
