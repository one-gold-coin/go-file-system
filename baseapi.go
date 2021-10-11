package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
)

//转发
func reverseProxy(ctx *gin.Context, filePath string) {
	proxy := httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL.Scheme = "http"
			request.URL.Host = FlagStaticProxyHost
			request.URL.Path = "/" + FlagStaticPath + "/" + filePath
		},
	}
	proxy.ServeHTTP(ctx.Writer, ctx.Request)
	ctx.Abort()
}
