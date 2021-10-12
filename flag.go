package main

import (
	"flag"
)

const (
	// EnvFile 配置文件路径
	EnvFile = "./.env.example"
)

var (
	FlagEnvFile string //配置文件路径
)

func init() {
	flag.StringVar(&FlagEnvFile, "f", EnvFile, "env配置文件,默认:"+EnvFile)
	//从arguments中解析注册的flag
	flag.Parse()
}
