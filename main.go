package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/one-gold-coin/getenv"
)

// go run *.go -f ./file

func main() {
	router := gin.Default()

	//初始化配置文件
	env := getenv.GetEnv{}
	env.SetFilePath(FlagEnvFile).Init()

	//注册路由
	Routers(router)

	//启动服务
	err := router.Run(":" + getenv.GetVal("APP_PORT").String())
	if err != nil {
		fmt.Println(" server err: ", err.Error())
		return
	}
}
