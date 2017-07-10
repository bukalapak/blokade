package main

import (
	"flag"
	"log"
	"net/http"

	bp "github.com/bukalapak/blokade/proxy"
	"github.com/elazarl/goproxy"
)

func main() {
	verbose := flag.Bool("v", false, "verbose mode")
	address := flag.String("addr", "127.0.0.1:8080", "proxy listen address")
	ignoredPath := flag.String("ipath", "", "ignored path")

	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	proxy.OnRequest().HandleConnect(goproxy.AlwaysReject)
	proxy.OnRequest(goproxy.Not(bp.IsLocalhost())).Do(bp.NotFoundHandler())
	proxy.OnRequest(bp.IsIgnored(*ignoredPath)).Do(bp.NotFoundHandler())

	log.Fatal(http.ListenAndServe(*address, proxy))
}
