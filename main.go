package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	bp "github.com/subosito/blokade/proxy"
	"log"
	"net/http"
)

func main() {
	verbose := flag.Bool("v", false, "verbose mode")
	address := flag.String("addr", "127.0.0.1:8080", "proxy listen address")

	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	proxy.OnRequest().HandleConnect(goproxy.AlwaysReject)
	proxy.OnRequest(goproxy.Not(bp.IsLocalhost())).Do(bp.NotFoundHandler())

	log.Fatal(http.ListenAndServe(*address, proxy))
}
