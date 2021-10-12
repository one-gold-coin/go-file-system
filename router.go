package main

import (
	"github.com/gin-gonic/gin"
	"github.com/one-gold-coin/getenv"
)

func Routers(router *gin.Engine) {
	//上传示例
	if getenv.GetVal("APP_ENV").String() == "dev" {
		router.Static("/demo", "./index")
	}

	if getenv.GetVal("STATIC_PROXY_OPEN").Int() > 0 {
		router.Static("/"+getenv.GetVal("STATIC_PROXY_PATH").String(), getenv.GetVal("FILE_PATH").String())
	}

	//文件上传
	router.POST("/upload", upload)

	//访问文件支持所有文件
	//根据文件md5值判断文件路径
	//如果不存在则404
	//如果存在则转发
	router.GET("/file/:file_key", file)

	//原始图片访问
	router.GET("/img/original/:file_key", imgOriginal)

	//压缩图片访问
	router.GET("/img/scale/:file_key/:width/:height/:quality", imgScale)

}
