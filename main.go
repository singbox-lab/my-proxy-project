package main

import (
    "fmt"
    "log"

    "my-proxy-project/pkg/config"
    "my-proxy-project/pkg/proxy"
)

func main() {
    fmt.Println("MyProxy — A simple multi-protocol proxy framework")

    cfg, err := config.LoadConfig("examples/config.json")
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    proxy.StartSOCKS(cfg.SOCKS.Listen)
}
