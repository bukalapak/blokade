package proxy

import (
	"github.com/elazarl/goproxy"
	"net/http"
	"regexp"
	"strings"
)

func IsLocalhost() goproxy.ReqConditionFunc {
	return func(req *http.Request, ctx *goproxy.ProxyCtx) bool {
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
}
