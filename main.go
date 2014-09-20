package main

import (
	"flag"
	"github.com/elazarl/goproxy"
	"log"
	"net/http"
	"regexp"
	"strings"
)

var isLocal goproxy.ReqConditionFunc = func(req *http.Request, ctx *goproxy.ProxyCtx) bool {
	host := req.URL.Host

	if strings.Contains(host, ":") {
		hostPort := strings.Split(host, ":")
		host = hostPort[0]
	}

	return host == "::1" ||
		host == "0:0:0:0:0:0:0:1" ||
		host == "localhost" ||
		regexp.MustCompile(`127\.0\.0\.\d+`).MatchString(host) ||
		strings.Contains(host, "lvh.me") ||
		strings.Contains(host, "xip.io")
}

func main() {
	verbose := flag.Bool("v", false, "verbose mode")
	address := flag.String("addr", "127.0.0.1:8080", "proxy listen address")

	flag.Parse()

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = *verbose

	proxy.OnRequest().HandleConnect(goproxy.AlwaysReject)
	proxy.OnRequest().DoFunc(
		func(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
			if isLocal(req, ctx) {
				return req, nil
			} else {
				return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusNotFound, "NOT FOUND!")
			}
		},
	)

	log.Fatal(http.ListenAndServe(*address, proxy))
}
