package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// go run *.go -FilePath ./file -Port 9081

func main() {
	router := gin.Default()

	//初始化参数
	initFlag()


	//上传示例
	if FlagAppEnv == "dev" {
		router.Static("/demo", "./demo")
	}

	router.Static("/"+FlagStaticPath, FlagFilePath)

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

	//启动服务
	err := router.Run(":" + strconv.Itoa(FlagPort))
	if err != nil {
		fmt.Println(" server err: ", err.Error())
		return
	}
}
