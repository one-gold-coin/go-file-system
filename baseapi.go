package main

import (
	"github.com/gin-gonic/gin"
	"github.com/one-gold-coin/getenv"
	"net/http"
	"net/http/httputil"
)

//转发
func reverseProxy(ctx *gin.Context, filePath string) {
	proxy := httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL.Scheme = getenv.GetVal("STATIC_PROXY_SCHEME").String()
			request.URL.Host = getenv.GetVal("STATIC_PROXY_HOST").String()
			request.URL.Path = "/" + getenv.GetVal("STATIC_PROXY_PATH").String() + "/" + filePath
		},
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
	ctx.Abort()
}
