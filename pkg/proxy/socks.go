package proxy

import (
    "log"

    "github.com/armon/go-socks5"
)

func StartSOCKS(listen string) {
    conf := &socks5.Config{}
    server, err := socks5.New(conf)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Starting SOCKS5 server on", listen)
    if err := server.ListenAndServe("tcp", listen); err != nil {
        log.Fatal(err)
    }
}
