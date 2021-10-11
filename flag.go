package main

import (
	"flag"
	"strconv"
)

func initFlag() {
	flag.StringVar(&FlagAppEnv, "AppEnv", AppEnv, "系统运行环境,默认:"+AppEnv)
	flag.IntVar(&FlagPort, "Port", Port, "端口,默认:"+strconv.Itoa(Port))
	flag.StringVar(&FlagFilePath, "FilePath", FilePath, "文件存储路径,默认:"+FilePath)
	flag.StringVar(&FlagFileType, "FileType", FileType, "可上传文件类型,默认:"+FileType)
	flag.StringVar(&FlagStaticPath, "StaticPath", StaticPath, "文件访问path,默认:"+StaticPath)
	flag.StringVar(&FlagStaticProxyHost, "StaticProxyHost", StaticProxyHost, "图片访问代理,默认:"+StaticProxyHost)
	//从arguments中解析注册的flag
	flag.Parse()
}
