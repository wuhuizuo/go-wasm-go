wasmedge_version := "0.9.0"
tinygo_ver := "0.21.0"

install_wasmedge_shared_lib:
    wget -qO- https://raw.githubusercontent.com/WasmEdge/WasmEdge/master/utils/install.sh | bash -s -- -v {{wasmedge_version}} -e all -p /usr/local/ -r yes

install_tinygo:    
    wget -qO- https://github.com/tinygo-org/tinygo/releases/download/v{{tinygo_ver}}/tinygo{{tinygo_ver}}.linux-amd64.tar.gz | tar -zxf - -C /usr/local/
    echo 'PATH=$PATH:/usr/local/tinygo/bin' > /etc/profile.d/tinygo.sh

build_wasm_tinygo:
    #!/usr/bin/env sh
    cd provider/wasm-tinygo &&
    tinygo build -target=wasi -o wasm.wasm

build_wasm_go:
    #!/usr/bin/env sh
    cd provider/wasm-go
    GOOS=js GOARCH=wasm go build -o wasm.wasm

build_plugin:
    #!/usr/bin/env sh
    cd provider/plugin
    go build -buildmode=plugin -o plugin.so